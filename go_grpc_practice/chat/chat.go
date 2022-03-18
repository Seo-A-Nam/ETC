package chat

import (
	"log" // log every incoming requests
	"golang.org/x/net/context" // get the context for the request
)

type Server struct {

}
 
func (s *Server) SayHello(ctx context.Context, message *Message) (*Message,
error) { // define rpc methods - returns message pointer or error
	log.Printf("Recieved message body from client: %s", message.body)
	return &Message{Body: "Hello From the Server!"}, nil // if fail, return nil for error
}
