package main

import (
	"context"
	pb "grpc-go/blog/proto"
	"log"
)

func createBlog(c pb.BlogServiceClient) string {
	log.Println("createBlog was invoked")

	res, err := c.CreateBlog(context.Background(), &pb.Blog{
		AuthorId: "1",
		Title:    "title",
		Content:  "content",
	})

	if err != nil {
		log.Fatalf("unexpected error: %v\n", err)
	}

	log.Printf("Blog has been created: %v\n", res.Id)
	return res.Id
}
