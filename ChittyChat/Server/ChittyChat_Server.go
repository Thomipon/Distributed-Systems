package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"sync"

	"main/clock"
	pb "main/proto"

	"google.golang.org/grpc"
)

type message_stream = pb.ChittyChatService_EnterChatServer

type server struct {
	pb.UnimplementedChittyChatServiceServer
	clients map[string]message_stream
	mutex   sync.Mutex // This could be split into read/write locks or by resource but this should suffice for our case
	clock   clock.LamportClock
}

type message_error struct {
	msg string
}

func (e *message_error) Error() string {
	return e.msg
}

var (
	port = flag.Int("port", 50051, "The server port")
)

func broadcast_internal(s *server, message string, time uint64) {
	log.Printf("Broadcasting message \"%s\" at time %d\n", message, time)
	for _, stream := range s.clients {
		stream.Send(&pb.MessageReply{Message: message, Time: time})
	}
}

func register_client(s *server, name string, stream message_stream) bool {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	_, exists := s.clients[name]
	if exists {
		return false
	}
	s.clients[name] = stream
	time := s.clock.Tick()
	broadcast_internal(s, fmt.Sprintf("Participant %s joined Chitty-Chat at Lamport time %d", name, time), time)
	return true
}

func deregister_client(s *server, name string) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	time := s.clock.Tick()
	broadcast_internal(s, fmt.Sprintf("Participant %s left Chitty-Chat at Lamport time %d", name, time), time)
	delete(s.clients, name)
}

func broadcast(s *server, message string, time uint64) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	broadcast_internal(s, message, time)
}

func (s *server) EnterChat(stream message_stream) error {
	msg, err := stream.Recv()
	if err != nil {
		return err
	}
	s.clock.Advance(msg.GetTime())
	name := msg.GetMessage() // first message is the requested name
	if !register_client(s, name, stream) {
		log.Println("Someone tried to log in as %s but this name is already taken by another user!", name)
		return &message_error{"Name is already taken"}
	}
	defer deregister_client(s, name)
	for {
		msg, err := stream.Recv()
		if err != nil {
			break
		}
		final_message := fmt.Sprintf("%s: %s", name, msg.GetMessage())
		s.clock.Advance(msg.GetTime()) // receive
		time := s.clock.Tick()         // send
		broadcast(s, final_message, time)
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
	s := server{}
	s.clients = make(map[string]message_stream)
	pb.RegisterChittyChatServiceServer(serv, &s)
	log.Printf("server listening at %v", listen.Addr())
	if err := serv.Serve(listen); err != nil {
		log.Fatalf("Failed to serve with error: %v", err)
	}
}
