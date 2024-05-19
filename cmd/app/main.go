package main

import (
	"context"
	"flag"
	"log"
	"time"
	"tourism/internal/app"
)

var (
	configPath string
)

// @title Medical Chat App API
// @version 1.0
// @description  API Server for Medical Chat App
// @host localhost:8080
// @BasePath /api/v1

func main() {
	flag.StringVar(&configPath, "config-path", "config/config.yaml", "path to configuration file")
	flag.Parse()

	appl, err := app.New(configPath)
	if err != nil {
		log.Fatalf("Initializing app: %v\n", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	if err = appl.Run(ctx); err != nil {
		log.Fatalf("Running an application: %v\n", err)
	}
}
