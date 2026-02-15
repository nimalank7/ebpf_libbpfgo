#include <linux/types.h>
#include <linux/bpf.h>
#include <bpf/bpf_helpers.h>
#include <bpf/bpf_tracing.h>
#include <bpf/bpf_core_read.h>

typedef __u32 u32;
typedef __u64 u64;

#define BPF_PERF_OUTPUT(_name)                          \
    struct {                                            \
        __uint(type, BPF_MAP_TYPE_PERF_EVENT_ARRAY);    \
        __uint(key_size, sizeof(u32));                  \
        __uint(value_size, sizeof(u32));                \
    } _name SEC(".maps");                               \

char LICENSE[] SEC("license") = "Dual BSD/GPL";