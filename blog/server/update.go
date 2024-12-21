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

func (s *Server) UpdateBlog(ctx context.Context, in *pb.Blog) (*emptypb.Empty, error) {
	log.Printf("UpdateBlog was invoked with %v\n", in)

	old, err := primitive.ObjectIDFromHex(in.Id)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "Cannot parse ID")
	}

	data := &BlogItem{
		AuthorId: in.AuthorId,
		Title:    in.Title,
		Content:  in.Content,
	}

	res, err := collection.UpdateOne(ctx, bson.M{"_id": old}, bson.M{"$set": data})

	if err != nil {
		return nil, status.Errorf(codes.Internal, fmt.Sprintf("Could not update: %v\n", err))
	}

	if res.MatchedCount == 0 {
		return nil, status.Errorf(codes.NotFound, "cannot find blog with id")
	}

	return &emptypb.Empty{}, nil
}
