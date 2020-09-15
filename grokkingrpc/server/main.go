package main

import (
	"context"
	"log"
	"net"

	pb "github.com/senyosimpson/tutorials/grokkingrpc/helloworld"
	"google.golang.org/grpc"
)

type helloWorldServer struct{}

func (s *helloWorldServer) Greet(ctx context.Context, request *pb.HelloRequest) (*pb.HelloReply, error) {
	log.Printf("Received: %v", request.GetName())
	return &pb.HelloReply{Message: "Hello " + request.GetName()}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	server := grpc.NewServer()
	pb.RegisterHelloWorldServer(server, &helloWorldServer{})
	if err := server.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
