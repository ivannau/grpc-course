package main

import (
	"fmt"
	"log"

	"github.com/ivannau/grpc-course/greet/greetpb"
	"google.golang.org/grpc"
)

func main() {

	fmt.Println("Hello from client")

	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())

	if err != nil {
		log.Fatalf("Could not connect: %v", err)
	}

	defer cc.Close()

	c := greetpb.NewGreetServiceClient(cc)

	fmt.Printf("created cleint: %f", c)
}
