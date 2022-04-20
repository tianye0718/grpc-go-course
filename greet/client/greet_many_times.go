package main

import (
	"context"
	"io"
	"log"

	pb "github.com/tianye0718/grpc-go-course/greet/proto"
)

func doGreetManyTimes(c pb.GreetServiceClient) {
	log.Println("doGreetManyTimes was involved")

	req := &pb.GreetRequest{
		FirstName: "Ye",
	}

	stream, err := c.GreetManyTimes(context.Background(), req)
	if err != nil {
		log.Fatalf("Error while calling GreetManyTimes: %v\n", err)
	}

	for {
		msg, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("Error while reading the stream: %v", err)
		}

		log.Printf("Received message from server: %s\n", msg)
	}
}
