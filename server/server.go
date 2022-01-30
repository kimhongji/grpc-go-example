package main

import (
	pb "grpc-go-example/api/proto"
	"grpc-go-example/server/handler"
	"log"
	"net"

	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", "localhost:50051")
	if err != nil {
		log.Fatalf("failed to listen %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterApiServer(grpcServer, &handler.APIServer{})

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}