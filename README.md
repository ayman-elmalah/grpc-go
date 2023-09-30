# gRPC Go Project

This is a gRPC project in Go that demonstrates four different gRPC API types, each in its own project directory. You can clone this repository to get started with the project.

The project is organized into four separate directories, each representing a different gRPC API type:

1. **Unary API:** Demonstrates a simple unary gRPC API where a client sends a request and receives a single response from the server.

    - Directory: `go-grpc-unary/`

2. **Server Streaming API:** Shows an example of a server streaming gRPC API where a client sends a request and receives a stream of responses from the server.

    - Directory: `go-grpc-server-streaming/`

3. **Client Streaming API:** Illustrates a client streaming gRPC API where a client sends a stream of requests to the server and receives a single response.

    - Directory: `go-grpc-server-client-streaming/`

4. **Bidirectional Streaming API:** Demonstrates a bidirectional streaming gRPC API where both the client and server can send a stream of messages to each other.

    - Directory: `go-grpc-server-bidirectional-streaming/`

Each project directory contains its own Go code, including the necessary `.proto` files for defining the gRPC service, server and client implementations, and Makefile (if applicable).

## Getting Started

To get started with a specific gRPC API type, navigate to the corresponding project directory and follow the instructions in the README file provided in that directory.

## Installation

To get started, clone this repository to your local machine using the following command:

```bash
git clone https://github.com/ayman-elmalah/grpc-go
```

## License

This project is licensed under the MIT License.