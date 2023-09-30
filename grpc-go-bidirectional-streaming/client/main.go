package main

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"grpc-go-bidirectional-streaming/proto"
	"log"
	"time"
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

	done := make(chan bool)

	// Send requests in a separate goroutine
	go func() {
		defer close(done) // Ensure we close the 'done' channel when done
		for _, object := range objects {
			request := &proto.HiRequest{
				Name: object.Name,
				Age:  object.Age,
			}
			if err := stream.Send(request); err != nil {
				log.Fatalf("Error sending request: %v", err)
			}
			time.Sleep(1 * time.Second)
		}
		stream.CloseSend()
	}()

	// Receive responses in the main goroutine
	for {
		response, err := stream.Recv()
		if err != nil {
			break // Exit the loop when the stream is closed
		}
		log.Printf("Received response: %s", response.GetMessage())
	}

	<-done // Wait for the 'done' channel to close, indicating that the goroutine has finished

}
