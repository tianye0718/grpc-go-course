package main

import (
	"context"
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
