package database

import (
	"log/slog"
	"os"

	"github.com/ajaysinghnp/dcabinet/lib/config"
	"github.com/ajaysinghnp/dcabinet/lib/database/sqlite"
)

func MustSetup(config *config.Config) Database {
	db, err := sqlite.New(config)
	if err != nil {
		slog.Error("Failed to initialize database", "error", err)
		os.Exit(0)
	}

	// Initialize the database schema if needed
	if err := db.InitializeTables(); err != nil {
		slog.Error("Failed to initialize database tables", "error", err)
		os.Exit(0)
	}

	slog.Info("Database Initialized Successfully")

	return db
}
