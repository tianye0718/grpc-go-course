package main

import (
	"fmt"
	"io"
	"log"

	pb "github.com/tianye0718/grpc-go-course/greet/proto"
)

func (s *Server) LongGreet(stream pb.GreetService_LongGreetServer) error {
	log.Println("LongGreet service was involved")
	res := ""

	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&pb.GreetResponse{
				Result: res,
			})
		}
		if err != nil {
			log.Fatalf("Error while reading client stream: %v", err)
		}
		res += fmt.Sprintf("hello %s!\n", req.FirstName)
	}
}
