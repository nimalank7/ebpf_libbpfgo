# Repository guidance

## Project structure

- `cmd/` contains the Go eBPF examples. Each example's Go loader is paired with a `.bpf.c` program.
- `libbpf/` contains the libbpf source used to build the examples.
- `internal/` contains shared Go code.

## Changes

- Keep changes to an example's Go loader and its eBPF C program consistent.
- Format Go changes with `gofmt`.
- Preserve the existing libbpfgo and libbpf integration; avoid changing vendored or submodule code unless the task specifically requires it.
- eBPF builds and runtime behavior depend on Linux tooling and kernel support. For build instructions, use the Docker and Vagrant workflow documented in `README.md`.

## Validation

- Run relevant Go tests with `go test ./...` when the environment has the required libbpf build dependencies.
- For changes to an eBPF program, build the corresponding example in the documented Linux/Docker environment; macOS alone does not provide the Linux eBPF toolchain or runtime.
