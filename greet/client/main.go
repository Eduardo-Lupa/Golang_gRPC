package main

import (
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	pb "github.com/Eduardo-Lupa/Golang_gRPC/greet/proto"
)

var adress string = "localhost:50051"

func main() {
	conn, err := grpc.Dial(adress, grpc.WithTransportCredentials(insecure.NewCredentials())) // passando o endereco de memoria e sem credenciais ainda

	if err != nil {
		log.Fatalf("Failed to connect: %v\n", err)
	}

	defer conn.Close()

	c := pb.NewGreetServiceClient(conn)
	doGreet(c)
}