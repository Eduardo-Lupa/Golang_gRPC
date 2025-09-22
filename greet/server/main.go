package main

import (
	"log"
	"net"

	pb "github.com/Eduardo-Lupa/Golang_gRPC/greet/proto"
	"google.golang.org/grpc"
)

var address string = "0.0.0.0:50051" // localhost

// struct to all RPC endpoints
type Server struct {
	pb.GreetServiceServer
}

func main() {
	lis, err := net.Listen("tcp", address) // 
	if err != nil {
		log.Fatalf("failed to listen on: %v\n", err)
	}

	log.Printf("listening on: %v\n", address)

	s := grpc.NewServer()
	pb.RegisterGreetServiceServer(s, &Server{})

	if err = s.Serve(lis); err != nil {
		log.Fatalf("failed to server on: %v\n", err)
	}
	
}