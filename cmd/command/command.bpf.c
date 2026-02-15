// +build ignore
#include "../../hello.bpf.h"

// Define a perf buffer called gotopia
BPF_PERF_OUTPUT(gotopia);
// kprobe on execve syscall
SEC("kprobe/sys_execve")
int hello(void *ctx) {
    // Define an array
    char data[30];
    // Retrieve the command and write the information to data
    bpf_get_current_comm(&data, sizeof(data));
    // Send data to the gotopia perf ring buffer
    bpf_perf_event_output(ctx, &gotopia, BPF_F_CURRENT_CPU, &data, sizeof(data));
    return 0;
}