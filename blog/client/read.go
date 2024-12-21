package main

import (
	"context"
	pb "grpc-go/blog/proto"
	"log"
)

func readBlog(c pb.BlogServiceClient, id string) *pb.Blog {
	log.Println("readBlog was invoked")

	res, err := c.ReadBlog(context.Background(), &pb.BlogId{
		Id: id,
	})

	if err != nil {
		log.Fatalf("unexpected error: %v\n", err)
	}

	log.Printf("Blog has been red: %v\n", res)
	return res
}
