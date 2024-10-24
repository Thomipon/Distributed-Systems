package main

import (
	"bufio"
	"context"
	"flag"
	"log"
	"os"
	"strings"

	clock "main/Clock"
	pb "main/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type message_stream = pb.ChittyChatService_EnterChatClient

var (
	addr = flag.String("addr", "localhost:50051", "the address to connect to")
)

func read(reader *bufio.Reader) string {
	text, _ := reader.ReadString('\n')
	text = strings.Replace(text, "\r\n", "", -1)
	return strings.Replace(text, "\n", "", -1)
}

func send(stream message_stream, clock *clock.LamportClock) {
	reader := bufio.NewReader(os.Stdin)

	name_entered := false

	for {
		if !name_entered {
			log.Println("Please enter your desired name:")
		}
		text := read(reader)
		if len(text) == 0 {
			continue
		}
		if text[0] == '\\' {
			// Command
			if text == "\\exit" {
				log.Println("Exiting...")
				return
			}
			log.Printf("Unknown command: %s", text)
			continue
		}
		time := clock.Tick()
		err := stream.Send(&pb.MessageRequest{Message: text, Time: time})
		if err != nil {
			log.Fatalf("Error occurred while sending: %v", err)
			continue
		}
		name_entered = true // First message is the name; we only ever get here when entering the name is successful. This is a bit hacky and there are problems. Ideally we would want to be able to retry but right now we just crash
	}

}

func receive(stream message_stream, clock *clock.LamportClock) {
	for {
		msg, err := stream.Recv()
		if err != nil {
			log.Fatalf("Error occurred while receiving: %v", err)
			continue // TODO: This might be the wrong behaviour here; it seems to cause errors when exiting normally
		}
		time := clock.Advance(msg.GetTime())
		log.Printf("Received \"%s\" at time %d\n", msg.GetMessage(), time)
	}

}

func main() {
	flag.Parse()
	// Set up a connection
	connection, err := grpc.NewClient(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect with error: %v", err)
	}
	defer connection.Close()
	client := pb.NewChittyChatServiceClient(connection)

	// Contact the server without timeout
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	stream, err := client.EnterChat(ctx)
	if err != nil {
		log.Fatalf("Failed to enter chat with error: %v", err)
	}
	clock := clock.LamportClock{}
	go receive(stream, &clock) // TODO: This probably wants to be somewhere else. We should probably only receive messages when we are logged in with our name
	send(stream, &clock)
}
