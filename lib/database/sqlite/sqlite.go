package sqlite

import (
	"database/sql"

	"github.com/ajaysinghnp/dcabinet/lib/config"

	_ "github.com/mattn/go-sqlite3" // Import the SQLite driver wihout using it directly
)

func New(config *config.Config) (*SQLite, error) {
	// Initialize SQLite connection here
	// For example, open a database connection using sql.Open
	db, err := sql.Open("sqlite3", config.Database)
	if err != nil {
		return nil, err
	}

	// Return a new SQLite instance
	return &SQLite{
		DB: db,
	}, nil
}
