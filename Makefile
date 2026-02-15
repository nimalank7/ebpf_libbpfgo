ARCH=$(shell uname -m)

TARGET := hello
TARGET_DIR := ./cmd/${TARGET}/${TARGET}
TARGET_BPF := ${TARGET_DIR}.bpf.o

GO_SRC := ${TARGET_DIR}.go
BPF_SRC := ${TARGET_DIR}.bpf.c

LIBBPF_HEADERS := /app/libbpf/src/root/usr/include
LIBBPF_OBJ := /app/libbpf/src/build/libbpf.a -lelf -lz

.PHONY: all
all: libbpf-compile $(TARGET) $(TARGET_BPF) 

libbpf-compile:
	@echo "Compiling libbpf"
	cd ./libbpf/src && mkdir -p ./libbpf/src/build ./libbpf/src/root
	$(MAKE) -C ./libbpf/src BUILD_STATIC_ONLY=y OBJDIR=build DESTDIR=root install install_uapi_headers
	cp ./libbpf/include/uapi/linux/*.h ./libbpf/src/root/usr/include/linux/

go_env := CC=clang CGO_CFLAGS="-I $(LIBBPF_HEADERS)" CGO_LDFLAGS="$(LIBBPF_OBJ)"
$(TARGET): $(GO_SRC)
	@echo "Building from: $(GO_SRC)"
	$(go_env) go build -o ./cmd/${TARGET} ${GO_SRC} 

# Create /target folder to drop Go binaries
$(TARGET_BPF): $(BPF_SRC)
	clang \
		-I /usr/include/$(ARCH)-linux-gnu \
		-I $(LIBBPF_HEADERS) \
		-O2 -g -c -target bpf \
		-o $@ $<

# Need to fix this
.PHONY: clean
clean:
	go clean

