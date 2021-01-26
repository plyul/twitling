package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
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
		log.Fatalf("error connecting to server: %v", err)
	}
	defer func() {
		if err := conn.Close(); err != nil {
			log.Fatalf("error closing server connection: %v", err)
		}
	}()
	client := api.NewMessagingServiceClient(conn)
	response, err := client.Ping(context.Background(), &api.PingRequest{Challenge: "ping"})
	if err != nil {
		log.Fatalf("error calling Ping method: %v", err)
	}
	fmt.Printf("Got response: %s", response.Response)
}
