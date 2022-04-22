package main

import (
	"context"
	"io"
	"log"
	"time"

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

func doAvg(c pb.CalculatorServiceClient) {
	log.Println("doAvg was involved")

	reqs := []*pb.AvgRequest{
		{Num: 1.0},
		{Num: 2.0},
		{Num: 3.0},
		{Num: 4.0},
	}

	stream, err := c.Avg(context.Background())
	if err != nil {
		log.Fatalf("Error while calling Avg service")
	}

	for _, req := range reqs {
		log.Printf("sending num to server: %f\n", req.Num)
		stream.Send(req)
		time.Sleep(time.Second)
	}

	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("Error while receiving result from server: %v\n", err)
	}
	log.Printf("Avg result received from server: %f\n", res.Result)
}
