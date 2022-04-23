package main

import (
	"context"
	"io"
	"log"
	"math"

	pb "github.com/tianye0718/grpc-go-course/calculator/proto"
)

func (s *Server) Sum(ctx context.Context, in *pb.SumRequest) (*pb.SumResponse, error) {
	log.Printf("Involving Sum service with %v\n", in)
	return &pb.SumResponse{
		Result: in.Num1 + in.Num2,
	}, nil
}

func (s *Server) Primes(in *pb.PrimesRequest, stream pb.CalculatorService_PrimesServer) error {
	log.Printf("Involving Sum service with %v\n", in)
	k := int32(2)
	n := in.Number
	for n > 1 {
		if n%k == 0 {
			stream.Send(&pb.PrimesResponse{Result: k})
			n = n / k
		} else {
			k++
		}
	}
	return nil
}

func (s *Server) Avg(stream pb.CalculatorService_AvgServer) error {
	log.Println("Avg was involved")
	sum := 0.0
	count := 0.0
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			if count == 0 {
				return stream.SendAndClose(&pb.AvgResponse{Result: 0.0})
			} else {
				return stream.SendAndClose(&pb.AvgResponse{Result: sum / count})
			}
		}
		if err != nil {
			log.Fatalf("Error while recieving numbers from client: %v\n", err)
		}
		count++
		sum += req.Num
		log.Printf("Received number from client: %f\n", req.Num)
	}
}

func (s *Server) Max(stream pb.CalculatorService_MaxServer) error {
	log.Println("Max was invoked")

	currentMax := math.MinInt32
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			log.Fatalf("Error while receiving client request: %v\n", err)
		}
		if req.Num > int32(currentMax) {
			currentMax = int(req.Num)
		}
		err = stream.Send(&pb.MaxResponse{Max: int32(currentMax)})
		if err != nil {
			log.Fatalf("Error while sending data to client: %v\n", err)
		}
	}
}
