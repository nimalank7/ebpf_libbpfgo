package main

import "C"
import (
	"fmt"

	"os"
	"os/signal"

	sys "main/internal/sys"

	bpf "github.com/aquasecurity/libbpfgo"
)

func main() {
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt)

	bpfFilePath := "/home/vagrant/shared/cmd/command/command.bpf.o"

	// Read the object file
	b, err := bpf.NewModuleFromFile(bpfFilePath)
	must(err)
	defer b.Close()

	// Load into the kernel
	must(b.BPFLoadObject())

	// Retrieve the program
	p, err := b.GetProgram("hello")
	must(err)

	// Attach the program to the kprobe for execve
	_, err = p.AttachKprobe(sys.Sys_execve)
	must(err)

	// Setup perf buffer to receive these events
	e := make(chan []byte, 300)
	pb, err := b.InitPerfBuf("gotopia", e, nil, 1024)
	must(err)
	// Deprecated use the new one
	pb.Start()

	// Loop over the buffer and we know the ELF
	go func() {
		for {
			data := <-e
			comm := string(data)
			fmt.Printf("Got %s\n", comm)
		}
	}()

	// Receive notification and print message
	<-sig
	fmt.Println("Cleaning up")
	//TODO - Investigate why this throws a panic
	pb.Stop()

}

// Convenience function to trap errors and just panic
func must(err error) {
	if err != nil {
		panic(err)
	}
}
