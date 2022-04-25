package main

import (
	"context"
	"log"

	pb "github.com/tianye0718/grpc-go-course/blog/proto"
)

func deleteBlog(c pb.BlogServiceClient, id string) {
	log.Println("deleteBlog was invoked")

	_, err := c.DeleteBlog(context.Background(), &pb.BlogId{Id: id})
	if err != nil {
		log.Fatalf("error while deleting: %v\n", err)
	}

	log.Println("Blog has been deleted")
}
