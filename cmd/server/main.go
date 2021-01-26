package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
	"twitling/generated/api"
)

const (
	apiAddress = "localhost"
	apiPort = 1025
)

type server struct {
	api.UnimplementedMessagingServiceServer
}

func (s *server) Ping(_ context.Context, req *api.PingRequest) (*api.PingResponse, error) {
	log.Printf("api.Ping called with arg: %v", req)
	return &api.PingResponse{
		Response: req.Challenge,
	}, nil
}

func main() {
	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%d", apiAddress, apiPort))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)
	api.RegisterMessagingServiceServer(grpcServer, &server{})
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("gRPC server error: %v", err)
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
