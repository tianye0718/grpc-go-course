package main

import (
	"context"
	"log"

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
