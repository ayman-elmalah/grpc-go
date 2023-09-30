package main

import (
	"fmt"
	"google.golang.org/grpc"
	"grpc-go-client-streaming/proto"
	"io"
	"log"
	"net"
)

type Server struct {
	proto.HiServiceServer
}

func (s *Server) SayHi(stream proto.HiService_SayHiServer) error {
	log.Print("sayHi function was invoked\n")

	res := ""

	for {
		req, err := stream.Recv()

		if err == io.EOF {
			return stream.SendAndClose(&proto.HiResponse{
				Message: res,
			})
		}

		if err != nil {
			break
		}

		res += fmt.Sprintf("Hi, %s your age is %d !\n", req.GetName(), req.GetAge())
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
