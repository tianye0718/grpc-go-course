package main

import (
	"context"
	"fmt"
	"log"

	pb "github.com/tianye0718/grpc-go-course/blog/proto"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *Server) ListBlogs(in *emptypb.Empty, stream pb.BlogService_ListBlogsServer) error {
	log.Printf("ListBlogs was invoked with %v\n", in)

	cur, err := collection.Find(context.Background(), primitive.D{{}})
	defer cur.Close(context.Background())

	if err != nil {
		return status.Errorf(
			codes.Internal,
			fmt.Sprintf("Internal Error while finding in mongoDB: %v\n", err),
		)
	}

	for cur.Next(context.Background()) {
		data := &BlogItem{}
		err := cur.Decode(data)
		if err != nil {
			return status.Errorf(
				codes.Internal,
				fmt.Sprintf("Internal Error while decoding data from mongoDB: %v\n", err),
			)
		}
		stream.Send(documentToBlog(data))
	}

	if err = cur.Err(); err != nil {
		return status.Errorf(
			codes.Internal,
			fmt.Sprintf("Unkonwn Internal Error: %v\n", err),
		)
	}

	return nil
}
