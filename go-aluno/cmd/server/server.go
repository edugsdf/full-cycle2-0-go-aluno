package main

import (
	"log"
	"net"

	"google.golang.org/grpc"
)

func main() {

	lis, err := net.Listen("tcp", "localhost:50051")
	if err != nil {
		log.Fatal("Could not conect: %v", err)
	}

	grpcServer := grpc.NewServer()

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatal("Could not server: %v", err)
	}
}
