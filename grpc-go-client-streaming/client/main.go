package main

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"grpc-go-client-streaming/proto"
	"log"
)

type Object struct {
	Name string
	Age  int32
}

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	client := proto.NewHiServiceClient(conn)

	stream, err := client.SayHi(context.Background())
	if err != nil {
		log.Fatalf("Error creating stream: %v", err)
	}

	objects := []Object{
		{Name: "Ayman Elmalah", Age: 28},
		{Name: "The Ayman", Age: 27},
		{Name: "Ayman", Age: 26},
	}

	for _, object := range objects {
		request := &proto.HiRequest{
			Name: object.Name,
			Age:  object.Age,
		}
		if err := stream.Send(request); err != nil {
			log.Fatalf("Error sending request: %v", err)
		}

		log.Print("Sending request\n")
	}

	response, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("Error receiving response: %v", err)
	}

	log.Printf("Received response: \n%s", response.GetMessage())
}
