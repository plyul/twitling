package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"io"
	"log"
	"twitling/generated/api"
)

const (
	apiAddress = "localhost"
	apiPort = 1025
)

func main() {
	conn, err := grpc.Dial(fmt.Sprintf("%s:%d", apiAddress, apiPort), grpc.WithInsecure())
	if err != nil {
		log.Fatalf("ошибка при соединении с сервром: %v", err)
	}
	defer func() {
		if err := conn.Close(); err != nil {
			log.Fatalf("ошибка при закрытии соединения с сервером: %v", err)
		}
	}()
	client := api.NewMessagingAPIClient(conn)
	stream, err := client.Notification(context.Background())
	if err != nil {
		log.Fatalf("ошибка вызова метода Notification: %v", err)
	}
	go func() {
		for _, note := range []string{"денег нет", "держитесь", "конец"} {
			if err := stream.Send(&api.Note{Text: note}); err != nil {
				log.Println("ошибка при отправке донесения")
			}
		}
	}()
	for {
		note, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			break
		}
		log.Printf("Получил ответ на донесение: %s", note.Text)
		if note.Text == "ack: конец" {
			break
		}
	}
	fmt.Println("Все донесения были успешно отправлены")
}
