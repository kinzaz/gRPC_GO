package main

import (
	pb "grpc-go/greet/proto"
	"io"
	"log"

	"google.golang.org/grpc"
)

func (s *Server) GreetEveryone(stream grpc.BidiStreamingServer[pb.GreetRequest, pb.GreetResponse]) error {
	log.Println("GreetEveryone function was invoked with")

	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			log.Fatalf("Error while reading client stream: %v\n", err)
		}

		res := "hello " + req.FirstName + "!"

		err = stream.Send(&pb.GreetResponse{
			Result: res,
		})

		if err != nil {
			log.Fatalf("Error while sending data to client: %v", err)
		}
	}
}
