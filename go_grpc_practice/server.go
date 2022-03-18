// simple chat service using grpc
package main

import (
	"log"
	"net"
	"google.golang.org/gprc"
)

func main() {
	lis, err := net.Listen("tcp", ":9000") // lis means listener
	if err != nil {
		log.Fatalf("Failed to listen on port 9000: %v", err)
	}
	
	grpcServer := grpc.NewServer()

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to server gRPC server over port 9000: %v", err)
	}
} 