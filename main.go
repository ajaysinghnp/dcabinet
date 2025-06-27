package main

import (
	"context"
	"log"
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	"github.com/ajaysinghnp/dcabinet/api"
	"github.com/ajaysinghnp/dcabinet/lib/config"
)

func main() {
	// load config
	config := config.MustLoad()

	// database setup

	// setup router
	router := api.NewRouter()

	// setup server
	http_addr := config.HTTPServer.Host + ":" + config.HTTPServer.Port
	server := api.NewServer(http_addr, router)

	// start server
	done := make(chan os.Signal, 1)

	defer close(done)

	signal.Notify(done, os.Interrupt, syscall.SIGTERM, syscall.SIGINT)

	go func() {
		if err := server.ListenAndServe(); err != nil {
			log.Fatal("Failed to start server: ", err)
		}
	}()

	slog.Info("Starting server on", "address", http_addr)
	<-done

	slog.Info("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), config.HTTPServer.Timeout)

	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		slog.Error("Server shutdown failed", "error", err)
	}

	slog.Info("Server Stopped Successfully")
	os.Exit(0)
}
