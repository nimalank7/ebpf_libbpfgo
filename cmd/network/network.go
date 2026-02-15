package main

import "C"
import (
	"fmt"

	"os"
	"os/signal"

	bpf "github.com/aquasecurity/libbpfgo"
	"github.com/aquasecurity/libbpfgo/helpers"
)

/*
Attach to localhost
ping localhost -c 10
*/
const deviceName = "lo"

func main() {
	sig := make(chan os.Signal, 1)
	// Handle interrupts
	signal.Notify(sig, os.Interrupt)

	// eBPF bytecode of network.bpf.c
	bpfFilePath := "/home/vagrant/shared/cmd/network/network.bpf.o"

	// Read the object file
	b, err := bpf.NewModuleFromFile(bpfFilePath)
	must(err)
	defer b.Close()

	// Load the object file into the kernel
	must(b.BPFLoadObject())

	// Retrieve the program
	p, err := b.GetProgram("hello")
	must(err)

	// Attach the program to the kprobe to fire for execve
	_, err = p.AttachXDP(deviceName)
	must(err)

	// Print out whatever it receives from the debug tracing
	go helpers.TracePipeListen()

	// Receive notification and print clean
	<-sig
	fmt.Println("Cleaning up")

}

// Convenience function to trap errors and just panic
func must(err error) {
	if err != nil {
		panic(err)
	}
}
