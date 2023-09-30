package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"grpc-go-unary/proto"
	"log"
	"time"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	client := proto.NewHiServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	var name string = "Ayman"
	var age int32 = 28

	req := &proto.HiRequest{
		Name: name,
		Age:  age,
	}

	response, err := client.SayHi(ctx, req)
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	fmt.Printf("Server response: %s\n", response.GetMessage())
}
