package main

import (
	"context"
	"fmt"
	"log"

	pb "github.com/tianye0718/grpc-go-course/blog/proto"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) CreateBlog(ctx context.Context, in *pb.Blog) (*pb.BlogId, error) {
	log.Printf("CreatBlog was invoked with %v\n", in)
	// prepare data that needs send to mongoDB
	data := BlogItem{
		AuthorId: in.AuthorId,
		Title:    in.Title,
		Content:  in.Content,
	}
	// insert to mongoDB
	res, err := collection.InsertOne(ctx, data)
	if err != nil {
		return nil, status.Errorf(
			codes.Internal,
			fmt.Sprintf("Internal Error while inserting to mongoDB: %v\n", err),
		)
	}
	// Get oid
	oid, ok := res.InsertedID.(primitive.ObjectID)
	if !ok {
		return nil, status.Errorf(
			codes.Internal,
			"Error while converting to OID",
		)
	}
	// return oid
	return &pb.BlogId{Id: oid.Hex()}, nil
}
