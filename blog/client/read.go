package main

import (
	"context"
	"log"

	pb "github.com/tianye0718/grpc-go-course/blog/proto"
)

func readBlog(c pb.BlogServiceClient, id string) *pb.Blog {
	log.Println("readBlog was invoked")

	req := &pb.BlogId{Id: id}
	res, err := c.ReadBlog(context.Background(), req)
	if err != nil {
		log.Printf("Error happened while reading blog (id: %s) from server: %v\n", id, err)
	}

	log.Printf("blog was read: %v\n", res)
	return res
}
