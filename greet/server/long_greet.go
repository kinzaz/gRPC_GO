package main

import (
	"fmt"
	pb "grpc-go/greet/proto"
	"io"
	"log"

	"google.golang.org/grpc"
)

func (s *Server) LongGreet(stream grpc.ClientStreamingServer[pb.GreetRequest, pb.GreetResponse]) error {
	fmt.Println("LongGreet was invoked")

	res := ""

	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&pb.GreetResponse{
				Result: res,
			})

		}
		if err != nil {
			log.Fatalf("Error while reading client stream %v\n", err)
		}

		res += fmt.Sprintf("Hello %s!", req.FirstName)
	}
}
