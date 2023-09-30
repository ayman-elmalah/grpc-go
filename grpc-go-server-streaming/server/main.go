package main

import (
	"fmt"
	"google.golang.org/grpc"
	"grpc-go-server-streaming/proto"
	"log"
	"net"
)

type Server struct {
	proto.HiServiceServer
}

func (s *Server) SayHi(req *proto.HiRequest, stream proto.HiService_SayHiServer) error {
	name := req.GetName()
	age := req.GetAge()

	log.Printf("sayHi function was invoked with %v\n", req)

	for i := 0; i < 5; i++ {
		response := &proto.HiResponse{
			Message: fmt.Sprintf("Hi, %s your age is %d your number is %d!", name, age, i+1),
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
