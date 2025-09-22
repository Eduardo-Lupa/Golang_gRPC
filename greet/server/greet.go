package main

import (
	"context"
	"log"

	pb "github.com/Eduardo-Lupa/Golang_gRPC/greet/proto"
)

func (s *Server) Greet(ctx context.Context, in *pb.GreetRequest) (*pb.GreetResponse, error) {
	log.Printf("Greet function was invoked by: %v", in)
	return &pb.GreetResponse{
		Result: "Hii "+ in.FirstName,
	}, nil
}