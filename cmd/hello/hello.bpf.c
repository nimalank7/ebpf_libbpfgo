/*
This is the BPF code we want to run when execve is called
*/

// +build ignore
#include "../../hello.bpf.h"

SEC("kprobe/sys_execve")
int hello(void *ctx) {
	// eBPF helper function
	bpf_printk("Hello Gotopia!");
	return 0;
}
