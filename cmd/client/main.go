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
		log.Fatalf("Error connecting server: %v", err)
	}
	defer func() {
		if err := conn.Close(); err != nil {
			log.Fatalf("Error closing server connection: %v", err)
		}
	}()
	client := api.NewMessagingAPIClient(conn)
	stream, err := client.Notification(context.Background())
	if err != nil {
		log.Fatalf("Error calling Notification method: %v", err)
	}
	go func() {
		for _, note := range []string{"no money", "hold on", "end"} {
			if err := stream.Send(&api.Note{Text: note}); err != nil {
				log.Println("Error sending note")
			}
		}
	}()
	for {
		note, err := stream.Recv()
		if err == io.EOF {
			log.Println("Got EOF from server")
			break
		}
		if err != nil {
			log.Printf("Error receiving note: %v", err)
			break
		}
		log.Printf("Got notification reply: %s", note.Text)
		if note.Text == "ack: end" {
			fmt.Println("All notes were delivered successfully")
			break
		}
	}
}
