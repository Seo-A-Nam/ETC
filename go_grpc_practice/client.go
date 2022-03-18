package main

import (
	"log"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"github.com/Seo-A-Nam/Etc/go_grpc_practice/chat"
)

func main() {
	var	conn *grpc.ClientConn // connection pointer of grpc client
	conn, err := grpc.Dial(":9000", grpc.WithInsecure())
	// try to connect running grpc server on local host port(9000).
	// create insecure connection. -> disables cloud security for this cloud connection
	if err != nil {
		log.Fatalf("could not connect: %s", err)
	}
	defer conn.Close() // defer the connection close. 
	// cf) defer: execute this code line when function is gonna end.

	c := chat.NewChatServiceClient(conn) // passing the connection

	message := chat.Message {
		Body: "Hello from the client!",
	}

	response, err := c.SayHello(context.Background(), &message)
	if err != nil {
		log.Fatalf("Error when calling SayHello: %s", err)
	}
	log.Printf("Response from Server: %s", response.Body)
}