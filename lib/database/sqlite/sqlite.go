package sqlite

import (
	"database/sql"
	"log/slog"
	"os"

	"github.com/ajaysinghnp/dcabinet/lib/config"

	_ "github.com/mattn/go-sqlite3" // Import the SQLite driver wihout using it directly
)

func New(config *config.Config) (*SQLite, error) {
	// Initialize SQLite connection here
	// For example, open a database connection using sql.Open
	db, err := sql.Open(config.Database.Type, config.Database.URL)
	if err != nil {
		return nil, err
	}

	// Return a new SQLite instance
	return &SQLite{
		DB: db,
	}, nil
}

func (s *SQLite) Connect(config *config.Config) error {
	db, err := sql.Open(config.Database.Type, config.Database.URL)
	if err != nil {
		return err
	}
	s.DB = db
	return nil
}

func (s *SQLite) Close() error {
	// Close the database connection
	if s.DB != nil {
		return s.DB.Close()
	}
	return nil
}

func (s *SQLite) InitializeTables() error {
	// Create tables if they do not exist
	// Example: db.Exec("CREATE TABLE IF NOT EXISTS users (id INTEGER PRIMARY KEY, name TEXT)")
	// You can add more table creation logic here as needed
	s.initUsersTable()
	return nil
}

func (s *SQLite) initUsersTable() {
	// Create the users table if it does not exist
	query := `
	CREATE TABLE IF NOT EXISTS users (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		email TEXT NOT NULL UNIQUE
	);`
	_, err := s.DB.Exec(query)
	if err != nil {
		// Handle error, e.g., log it or return it
		slog.Error("Failed to create users table", "error", err)
		os.Exit(0)
	}
	slog.Info("Users table initialized successfully")
}
