package main

import "C"
import (
	"fmt"

	"os"
	"os/signal"

	sys "main/internal/sys"

	bpf "github.com/aquasecurity/libbpfgo"
	"github.com/aquasecurity/libbpfgo/helpers"
)

func main() {
	sig := make(chan os.Signal, 1)
	// Handle interrupts
	signal.Notify(sig, os.Interrupt)

	// eBPF bytecode of hello.bpf.c
	bpfFilePath := "/home/vagrant/shared/cmd/hello/hello.bpf.o"

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
	_, err = p.AttachKprobe(sys.Sys_execve)
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
