package main

import "C"
import (
	"fmt"

	"os"
	"os/signal"

	bpf "github.com/aquasecurity/libbpfgo"
)

func main() {
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt)

	bpfFilePath := "/home/vagrant/shared/cmd/counter_command/counter_command.bpf.o"

	// Read the object file
	b, err := bpf.NewModuleFromFile(bpfFilePath)
	must(err)
	defer b.Close()

	// Load into the kernel
	must(b.BPFLoadObject())

	// Retrieve the program
	p, err := b.GetProgram("hello")
	must(err)

	// Attach the program to the tracepoint for sys_enter
	_, err = p.AttachRawTracepoint("sys_enter")
	must(err)

	// Setup perf buffer to receive these events
	e := make(chan []byte, 300)
	pb, err := b.InitPerfBuf("gotopia", e, nil, 1024)
	must(err)
	pb.Start()

	// Create a map to hold the command names
	c := make(map[string]int, 1000)
	go func() {
		for {
			data := <-e
			comm := string(data)
			c[comm]++
		}
	}()

	// Receive notification and print message
	<-sig
	fmt.Println("Cleaning up")
	for comm, n := range c {
		fmt.Printf("%s: %d\n", comm, n)
	}

	//TODO - Investigate why this throws a panic
	pb.Stop()
}

// Convenience function to trap errors and just panic
func must(err error) {
	if err != nil {
		panic(err)
	}
}
