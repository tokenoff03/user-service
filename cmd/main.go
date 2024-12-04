package main

import (
	"context"
	"flag"
	"log"
	"user-service/internal/app"
)

var configPath string

func init() {
	flag.StringVar(&configPath, "config-path", "../.env", "path to config file")
}

func main() {
	flag.Parse()

	ctx := context.Background()
	app, err := app.NewApp(ctx, configPath)
	if err != nil {
		log.Fatalf("failed to init app: %v", err)
	}

	err = app.Run()
	if err != nil {
		log.Fatalf("failed to run app: %v", err)
	}
}
