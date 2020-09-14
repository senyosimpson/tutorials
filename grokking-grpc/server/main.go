package main

import (
	"context"
	pb "github.com/senyosimpson/tutorials/grokking-grpc/proto"
)

func Greet(ctx context.Context, request *pb.HelloRequest) (*pb.HelloReply, error) {
	log.Printf("Received: %v", in.GetName())
	return &pb.HelloReply{Message: "Hello " + in.GetName()}, nil
}