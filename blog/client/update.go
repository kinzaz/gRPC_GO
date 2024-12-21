package main

import (
	"context"
	pb "grpc-go/blog/proto"
	"log"
)

func updateBlog(c pb.BlogServiceClient, id string) {
	log.Println("updateBlog was invoked")

	res, err := c.UpdateBlog(context.Background(), &pb.Blog{
		Id:       id,
		AuthorId: "New Author id",
		Title:    "a new title",
		Content:  "a new content",
	})

	if err != nil {
		log.Fatalf("unexpected error: %v\n", err)
	}

	log.Printf("Blog has been updated: %v\n", res)
}
