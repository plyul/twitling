package main

import (
	"fmt"
	"google.golang.org/grpc"
	"io"
	"log"
	"net"
	"twitling/generated/api"
)

const (
	apiAddress = "localhost"
	apiPort = 1025
)

type server struct {
	api.UnimplementedMessagingAPIServer
}

func (s *server) Notification(stream api.MessagingAPI_NotificationServer) error {
	log.Println("Function api.Notification was started")
	for {
		note, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}
		log.Printf("Got note: %s", note.Text)
		if err := stream.Send(&api.Note{Text: fmt.Sprintf("ack: %s", note.Text)}); err != nil {
			log.Printf("Error sending reply: %v", err)
		}
	}
}

func main() {
	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%d", apiAddress, apiPort))
	if err != nil {
		log.Fatalf("Error establishing listener: %v", err)
	}
	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)
	api.RegisterMessagingAPIServer(grpcServer, &server{})
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Error registering gRPC server: %v", err)
	}
}

//func getUserName(userId int64) string {
//	for _, u := range state.Users {
//		if u.GetId() == userId {
//			return u.GetName()
//		}
//	}
//	return "Anonymous"
//}
