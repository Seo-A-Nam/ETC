// simple chat service using grpc
package main

import (
	"log"
	"net"
	"github.com/golang/protobuf/proto"
	"google.golang.org/gprc"
	"github.com/tutorialedge/go-grpc-tutorial/chat"
)

func main() {
	lis, err := net.Listen("tcp", ":9000") // lis means listener
	if err != nil {
		log.Fatalf("Failed to listen on port 9000: %v", err)
	}
	
	s := chat.Server() // define a new server

	grpcServer := grpc.NewServer() // make grpc server

	chat.RegisterChatServiceServer(grpcServer, &s) // register chat service server to grpc Server
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to server gRPC server over port 9000: %v", err)
	}
} 