package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"user-service/pkg/user_v1"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const grpcPort = 50051

type server struct {
	user_v1.UnimplementedUserV1Server
}

func (s *server) Get(ctx context.Context, req *user_v1.GetRequest) (*user_v1.GetResponse, error) {
	return &user_v1.GetResponse{}, nil
}

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", grpcPort))
	if err != nil {
		log.Fatalf("failed to listren: %v", err)
	}

	s := grpc.NewServer()
	reflection.Register(s)
	user_v1.RegisterUserV1Server(s, &server{})
	log.Printf("Server listening at %v", lis.Addr())

	if err = s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
