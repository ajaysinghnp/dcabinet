package main

import (
	"context"
	"log"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/ajaysinghnp/dcabinet/api"
	"github.com/ajaysinghnp/dcabinet/lib/config"
	"github.com/ajaysinghnp/dcabinet/lib/database/sqlite"
)

func main() {
	// load config
	config := config.MustLoad()
	slog.Info("Configuration Loaded")

	// database setup
	db, err := sqlite.New(config)
	if err != nil {
		slog.Error("Failed to initialize database", "error", err)
		os.Exit(0)
	}
	slog.Info("Database Initialized Successfully")

	defer db.DB.Close()

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
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal("Failed to start server: ", err)
		}
	}()

	slog.Info("Starting server on", "url", http_addr)

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
