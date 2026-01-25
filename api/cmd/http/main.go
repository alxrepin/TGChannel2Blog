// Package main API
//
//	@title			API
//	@version		1.0
//	@description	API
//	@host			localhost:8080
//	@BasePath		/api/v1
package main

import (
	"app/internal/bootstrap/app"
	"context"
	"log"
	"os/signal"
	"syscall"
	"time"
)
import "app/internal/bootstrap/config"

func main() {
	ctx, stop := signal.NotifyContext(
		context.Background(),
		syscall.SIGTERM,
		syscall.SIGINT,
	)
	defer stop()

	cfg := config.MustLoadConfig()
	application := app.New(ctx, cfg)

	go func() {
		<-ctx.Done()

		shutdown, cancel := context.WithTimeout(
			context.Background(),
			10*time.Second,
		)
		defer cancel()

		if err := application.Stop(shutdown); err != nil {
			log.Println("Graceful shutdown failed:", err)
		}
	}()

	log.Println("Application started")

	if err := application.Run(); err != nil {
		log.Fatal(err)
	}

	log.Println("Application stopped")
}
