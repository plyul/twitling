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
	log.Println("api.Notification начинает работу")
	for {
		note, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}
		log.Printf("Получил донесение: %s", note.Text)
		if err := stream.Send(&api.Note{Text: fmt.Sprintf("ack: %s", note.Text)}); err != nil {
			log.Printf("ошибка при отправке ответного донесения: %v", err)
		}
	}
}

func main() {
	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%d", apiAddress, apiPort))
	if err != nil {
		log.Fatalf("ошибка при попытке установки слушателя: %v", err)
	}
	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)
	api.RegisterMessagingAPIServer(grpcServer, &server{})
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("ошибка при запуске сервера gRPC: %v", err)
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
