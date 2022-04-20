package main

import (
	"context"
	"io"
	"log"

	pb "github.com/tianye0718/grpc-go-course/calculator/proto"
)

func doSum(c pb.CalculatorServiceClient) {
	log.Println("doSum was involved")
	res, err := c.Sum(context.Background(), &pb.SumRequest{
		Num1: 23,
		Num2: 12,
	})
	if err != nil {
		log.Fatalf("Could not sum: %v\n", err)
	}
	log.Printf("The result of Sum: %v\n", res.Result)
}

func doPrimes(c pb.CalculatorServiceClient) {
	log.Println("doPrimes was involved")

	req := &pb.PrimesRequest{
		Number: 120,
	}

	stream, err := c.Primes(context.Background(), req)
	if err != nil {
		log.Fatalf("Error while calling doPrimes: %v\n", err)
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
