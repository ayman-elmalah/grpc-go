# gRPC unary Hi Example

This is a simple gRPC project that demonstrates a client-server interaction where the client sends a name to the server, and the server responds with a greeting.

## Prerequisites

- Go should be installed on your system.
- `protoc` (Protocol Buffers compiler) should be installed. You can download it from https://github.com/protocolbuffers/protobuf/releases.

## Getting Started

1- Clone this repository to your local machine:
```bash
git clone https://github.com/ayman-elmalah/grpc-go.git && cd grpc-go/grpc-go-unary
```

2- Run `go mod tidy` to ensure the project's dependencies are up to date:
```bash
go mod tidy
```

3- Generate Go code from the hi.proto file:
```bash
make hi
```

4- Run the server start:
```bash
cd server && go run main.go
```

5- Run the client start:
- On another terminal tab, open in the grpc-go-unary directory then:
```bash
cd client && go run main.go
```