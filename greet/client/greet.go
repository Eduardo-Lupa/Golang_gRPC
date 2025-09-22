package main

import (
	"context"
	"log"

	pb "github.com/Eduardo-Lupa/Golang_gRPC/greet/proto"
)

func doGreet(c pb.GreetServiceClient) {
	log.Println("doGreet was invoked")
	res, err := c.Greet(context.Background(), &pb.GreetRequest{
		FirstName: "Eduardo",
	})
	if err != nil {
		log.Fatalf("Could not greet : %v \n", err)
	}
	log.Printf("Greeting: %v \n", res)
}
