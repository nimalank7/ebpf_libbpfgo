# eBPF Programs using libbpfgo

Repo contains 3 basic examples in Golang using [libbpfgo](https://github.com/aquasecurity/libbpfgo)
- hello
- command
- counter_command

## Build the Dockerfile

If using `arm64` then:

```
docker build --platform linux/arm64 -t hello .
```

If using `amd64` then:

```
docker build --platform linux/amd64 -t hello .
```

## Build the eBPF object and Go binary

To build the `hello` program run:

```
docker run -it -v $(pwd):/app hello TARGET=hello
```

Set `TARGET=counter` or `TARGET=counter_command` to build `counter` and `counter_command` programs respectively.

To read the ELF file run:

```
readelf -a hello.bpf.o
```

## Run the application

Vagrant uses VirtualBox and will map the current directory to `/home/vagrant/shared`

1. Launch Vagrant

```
vagrant up
cd shared
```

2. Run the program

```
sudo ./cmd/hello/hello
```

In another Vagrant session run some commands and wait a few seconds:

```
ls
echo "hello world"
```

3. Close the program by running `CNTRL + C`