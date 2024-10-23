package main

import (
	"flag"
	"fmt"
	"log"
	"net"

	pb "main/proto"

	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedChittyChatServiceServer
}

var (
	port = flag.Int("port", 50051, "The server port")
)

func (s *server) EnterChat(srv pb.ChittyChatService_EnterChatServer) error {
	for {
		msg, err := srv.Recv()
		if err != nil {
			break
		}
		log.Printf("Received: %v", msg.GetName())
		srv.Send(&pb.MessageReply{Message: "Hello " + msg.GetName()})
	}
	return nil
}

func main() {
	flag.Parse()
	listen, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("Failed to listen with error: %v", err)
	}
	serv := grpc.NewServer()
	pb.RegisterChittyChatServiceServer(serv, &server{})
	log.Printf("server listening at %v", listen.Addr())
	if err := serv.Serve(listen); err != nil {
		log.Fatalf("Failed to serve with error: %v", err)
	}
}
