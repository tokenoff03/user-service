package main

import (
	"context"
	"flag"
	"log"
	"net"
	"user-service/internal/config"
	"user-service/internal/config/env"
	userRepo "user-service/internal/repository/user"
	userService "user-service/internal/service/user"
	"user-service/pkg/user_v1"

	userAPI "user-service/internal/api/user"

	"github.com/jackc/pgx/v5/pgxpool"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configPath string

func init() {
	flag.StringVar(&configPath, "config-path", "../.env", "path to config file")
}

func main() {
	flag.Parse()

	ctx := context.Background()
	err := config.Load(configPath)
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	grpcConfig, err := env.NewGRPCConfig()
	if err != nil {
		log.Fatalf("failed to get grpc config: %v", err)
	}

	pgConfig, err := env.NewPGConfig()
	if err != nil {
		log.Fatalf("failed to get pg config: %v", err)
	}

	lis, err := net.Listen("tcp", grpcConfig.Address())
	if err != nil {
		log.Fatalf("failed to listren: %v", err)
	}

	pool, err := pgxpool.New(ctx, pgConfig.DSN())
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}
	defer pool.Close()

	userRepo := userRepo.NewRepository(pool)
	userService := userService.NewService(userRepo)
	s := grpc.NewServer()
	reflection.Register(s)
	user_v1.RegisterUserV1Server(s, userAPI.NewImplementation(userService))
	log.Printf("Server listening at %v", lis.Addr())

	if err = s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
