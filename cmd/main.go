package main

import (
	"context"
	"flag"
	"log"
	"user-service/internal/app"
)

var configPath string
var logLevel string

func init() {
	flag.StringVar(&configPath, "config-path", "../.env", "path to config file")
	flag.StringVar(&logLevel, "l", "info", "log level")
}

func main() {
	flag.Parse()

	ctx := context.Background()
	app, err := app.NewApp(ctx, configPath, logLevel)
	if err != nil {
		log.Fatalf("failed to init app: %v", err)
	}

	err = app.Run()
	if err != nil {
		log.Fatalf("failed to run app: %v", err)
	}
}
