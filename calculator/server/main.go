package main

import (
	pb "grpc-go/calculator/proto"
	"log"
	"net"

	"google.golang.org/grpc"
)

type Server struct {
	pb.CalculatorServiceServer
}

var addr string = "0.0.0.0:50051"

func main() {
	lis, err := net.Listen("tcp", addr)

	if err != nil {
		log.Fatalf("Failed to listen on %v\n", err)
	}
	log.Printf("Listening on %s\n", addr)

	var opts []grpc.ServerOption

	grpcServer := grpc.NewServer(opts...)
	pb.RegisterCalculatorServiceServer(grpcServer, &Server{})

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve %v\n", err)
	}
}
