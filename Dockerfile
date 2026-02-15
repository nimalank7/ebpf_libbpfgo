# Lower version as 20.10 didn't work
FROM ubuntu:latest

# Set Go version and architecture
ARG TARGETOS
ARG TARGETARCH
ARG GOVERSION=1.26.0

# Download development environment.
RUN apt-get update && \
    apt-get install -y \
        make \
        clang \
        llvm \
        libelf-dev

# Install Go specific version.
RUN apt-get install -y wget && \
    wget https://go.dev/dl/go${GOVERSION}.${TARGETOS}-${TARGETARCH}.tar.gz && \
    tar -xf go${GOVERSION}.${TARGETOS}-${TARGETARCH}.tar.gz && \
    mv go/ /usr/local/ && \
    ln -s /usr/local/go/bin/go /usr/local/bin/ && \
    rm -rf go${GOVERSION}.${TARGETOS}-${TARGETARCH}.tar.gz

# Setup working directory.
RUN mkdir -p /app
WORKDIR /app

# Build command
ENTRYPOINT ["/usr/bin/make", "all"]
# Override Makefile variable
CMD ["TARGET=hello"]