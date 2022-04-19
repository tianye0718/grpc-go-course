package main

import (
	"context"
	"log"

	pb "github.com/tianye0718/grpc-go-course/greet/proto"
)

func doGreet(c pb.GreetServiceClient) {
	log.Println("doGreet was involved")
	res, err := c.Greet(context.Background(), &pb.GreetRequest{
		FirstName: "Ye",
	})

	if err != nil {
		log.Fatalf("Could not greet: %v\n", err)
	}

	log.Printf("Greeting: %s\n", res.Result)
}
