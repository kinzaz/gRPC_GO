package main

import (
	"context"
	pb "grpc-go/blog/proto"
	"log"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *Server) ListBlogs(in *emptypb.Empty, stream grpc.ServerStreamingServer[pb.Blog]) error {
	log.Printf("ListBlog was invoked with %v\n", in)

	cur, err := collection.Find(context.Background(), primitive.D{})
	if err != nil {
		return status.Errorf(codes.Internal, "Unknown internal err")
	}
	defer cur.Close(context.Background())

	for cur.Next(context.Background()) {
		data := &BlogItem{}
		err := cur.Decode(data)

		if err != nil {
			return status.Errorf(codes.Internal, "Error while decoding data")
		}

		stream.Send(documentToBlog(data))
	}

	if err = cur.Err(); err != nil {
		return status.Errorf(codes.Internal, "Unknown internal err")
	}

	return nil
}
