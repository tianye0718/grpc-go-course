package main

import (
	"context"
	"io"
	"log"
	"time"

	pb "github.com/tianye0718/grpc-go-course/greet/proto"
)

func doGreetEveryone(c pb.GreetServiceClient) {
	log.Println("doGreetEveryone was invoked")

	stream, err := c.GreetEveryone(context.Background())
	if err != nil {
		log.Fatalf("Error while creating stream: %v\n", err)
	}

	reqs := []*pb.GreetRequest{
		{FirstName: "Ye"},
		{FirstName: "Vivian"},
		{FirstName: "Pier"},
	}

	waitc := make(chan struct{})

	go func() {
		for _, req := range reqs {
			log.Printf("Sending request: %v\n", req)
			stream.Send(req)
			time.Sleep(time.Second)
		}
		stream.CloseSend()
	}()

	go func() {
		for {
			res, err := stream.Recv()
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Printf("Error while receiving from server: %v\n", err)
				break
			}
			log.Printf("Received from server: %v\n", res.Result)
		}
		close(waitc)
	}()

	<-waitc
}
