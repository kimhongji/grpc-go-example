package main

import (
	"context"
	"log"
	"time"

	pb "grpc-go-example/api/proto"

	"google.golang.org/grpc"
)

const (
	address = "localhost:50051"
	name    = "AaronRoh"
)

func main() {
	// 서버 연결 셋업
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())

	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	c := pb.NewApiClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	reply, err := c.GetHello(ctx, &pb.Request{Name: name})
	// GetHello 호출
	if err != nil {
		log.Fatalf("GetHello error: %v", err)
	}
	log.Printf("Person: %v", reply)
}