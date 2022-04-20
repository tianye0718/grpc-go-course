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
