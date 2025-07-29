package main

import (
	"context"
	"log"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/lakshsetia/crud-docker/internal/config"
	"github.com/lakshsetia/crud-docker/internal/storage/postgresql"
)

func main() {
	// load config
	config, err := config.LoadConfig()
	if err != nil {
		log.Fatal(err)
	}

	// setup database
	postgresql, err := postgresql.New(config)
	if err != nil {
		log.Fatal(err)
	}

	// setup router
	router := http.NewServeMux()

	// setup handlers

	// setup server
	server := http.Server{
		Addr: config.HTTPServer.Address,
		Handler: router,
	}

	// start server
	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGTERM, syscall.SIGINT)
	slog.Info("starting server at ", slog.String("address", config.HTTPServer.Address))
	go func() {
		if err := server.ListenAndServe(); err != nil {
			log.Fatal("failed to start server: %w", err)
		}
	}()

	// shutdown server
	<-done
	ctx, cancel := context.WithTimeout(context.Background(), time.Second * 5)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		log.Fatal("failed to shutdown server: %w", err)
	}
	slog.Info("server shutdown successfully")
}