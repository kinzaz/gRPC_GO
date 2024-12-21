package main

import (
	"context"
	"fmt"
	pb "grpc-go/blog/proto"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *Server) DeleteBlog(ctx context.Context, in *pb.BlogId) (*emptypb.Empty, error) {
	log.Printf("DeleteBlog was invoked with %v\n", in)

	old, err := primitive.ObjectIDFromHex(in.Id)
	if err != nil {
		return nil, status.Errorf(codes.Internal, fmt.Sprintf("Cannot parse ID: %v\n", err))
	}

	res, err := collection.DeleteOne(ctx, bson.M{"_id": old})
	if err != nil {
		return nil, status.Errorf(codes.Internal, fmt.Sprintf("Cannot delete object in MongoDB: %v\n", err))
	}

	if res.DeletedCount == 0 {
		return nil, status.Errorf(codes.NotFound, "blog was not found")
	}

	return &emptypb.Empty{}, nil
}
