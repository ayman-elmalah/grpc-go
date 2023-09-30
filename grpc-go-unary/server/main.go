package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"grpc-go-unary/proto"
	"log"
	"net"
)

type Server struct {
	proto.HiServiceServer
}

func (s *Server) SayHi(ctx context.Context, req *proto.HiRequest) (*proto.HiResponse, error) {
	name := req.GetName()
	age := req.GetAge()

	log.Printf("sayHi function was invoked with %v\n", req)

	response := &proto.HiResponse{
		Message: fmt.Sprintf("Hi, %s your age is %d!", name, age),
	}

	return response, nil
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
