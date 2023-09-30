package main

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"grpc-go-server-streaming/proto"
	"log"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	client := proto.NewHiServiceClient(conn)

	var name string = "Ayman"
	var age int32 = 28

	req := &proto.HiRequest{
		Name: name,
		Age:  age,
	}

	stream, err := client.SayHi(context.Background(), req)
	if err != nil {
		log.Fatalf("Error calling SayHi: %v", err)
	}

	for {
		response, err := stream.Recv()
		if err != nil {
			log.Printf("Server stream closed: %v", err)
			break
		}
		log.Printf("Received: %s", response.GetMessage())
	}
}
