package main

import (
	"fmt"
	"google.golang.org/grpc"
	"grpc-go-bidirectional-streaming/proto"
	"log"
	"net"
)

type Server struct {
	proto.HiServiceServer
}

func (s *Server) SayHi(stream proto.HiService_SayHiServer) error {
	log.Print("sayHi function was invoked\n")

	for {
		req, err := stream.Recv()
		if err != nil {
			break
		}

		name := req.GetName()
		age := req.GetAge()

		response := &proto.HiResponse{
			Message: fmt.Sprintf("Hello, %s! Your age %d.\n", name, age),
		}
		if err := stream.Send(response); err != nil {
			return err
		}
	}
	return nil
}

func main() {
	listen, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	proto.RegisterHiServiceServer(s, &Server{})

	fmt.Println("Server is running on port 50051...")
	if err := s.Serve(listen); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
