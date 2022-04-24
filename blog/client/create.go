package main

import (
	"context"
	"log"

	pb "github.com/tianye0718/grpc-go-course/blog/proto"
)

func createBlog(c pb.BlogServiceClient) string {
	log.Println("createBlog was invoked")
	blog := &pb.Blog{
		AuthorId: "Ye",
		Title:    "My second blog",
		Content:  "content of my second blog",
	}

	res, err := c.CreateBlog(context.Background(), blog)
	if err != nil {
		log.Fatalf("Unexcepted error: %v\n", err)
	}

	log.Printf("Blog has been created with Id: %s\n", res.Id)
	return res.Id
}
