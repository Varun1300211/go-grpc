package main

import (
	"context"
	"log"

	pb "example.com/your-project-name/proto"
)

func callSayHelloServerStream(client pb.GreetServiceClient, names *pb.NamesList) {
	log.Printf("Streaming started")
	stream, err := client.SayHelloServerStreaming(context.Background(), names)
	if err != nil {
		log.Fatalf("Could not greet: %v", err)
	}

	for {
		message, err := stream.Recv()
		if err != nil {
			log.Printf("Error while receiving : %v", err)
			break
		}
		log.Println(message)
	}
	log.Printf("Streaming ended")
}
