package main

import (
	"context"
	"log"

	pb "github.com/tianye0718/grpc-go-course/greet/proto"
)

func (s *Server) Greet(ctx context.Context, in *pb.GreetRequest) (*pb.GreetResponse, error) {
	log.Printf("Involving Greet service with %v\n", in)
	return &pb.GreetResponse{
		Result: in.FirstName,
	}, nil
}
