package main

import (
	"context"
	"flag"
	"log"
	"time"

	pb "main/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// TODO: remove this
const (
	defaultName = "world"
)

var (
	addr = flag.String("addr", "localhost:50051", "the address to connect to")
	name = flag.String("name", defaultName, "Name to greet")
)

func main() {
	flag.Parse()
	// Set up a connection to the server.
	connection, err := grpc.NewClient(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect with error: %v", err)
	}
	defer connection.Close()
	client := pb.NewChittyChatServiceClient(connection)

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	stream, err := client.EnterChat(ctx)
	if err != nil {
		log.Fatalf("Failed to publish message with error: %v", err)
	}
	stream.Send(&pb.MessageRequest{Name: *name})
	msg, err := stream.Recv()
	if err == nil {
		log.Printf("Greeting: %s", msg.GetMessage())
	}
}
