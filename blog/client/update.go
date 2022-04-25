package main

import (
	"context"
	"log"

	pb "github.com/tianye0718/grpc-go-course/blog/proto"
)

func updateBlog(c pb.BlogServiceClient, id string) {
	log.Println("udpateBlog was invoked")

	newBlog := &pb.Blog{
		Id:       id,
		AuthorId: "Vivian",
		Title:    "A new title",
		Content:  "new content",
	}

	_, err := c.UpdateBlog(context.Background(), newBlog)
	if err != nil {
		log.Fatalf("Unexcepted error: %v\n", err)
	}

	log.Println("Blog has been updated")
}
