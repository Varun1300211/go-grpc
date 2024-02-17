package main

import (
	"log"
	"net"

	pb "example.com/your-project-name/proto"
	"google.golang.org/grpc"
)

const (
	port = ":8080"
)

type helloServer struct {
	pb.GreetServiceServer
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Failed to start the server %v", err)
	}
	grpcServer := grpc.NewServer()
	helloServer := &helloServer{}
	pb.RegisterGreetServiceServer(grpcServer, helloServer)
	log.Printf("Server started on port %v", lis.Addr())
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to start the server %v", err)
	}
}
