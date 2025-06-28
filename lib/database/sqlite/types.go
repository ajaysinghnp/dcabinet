package sqlite

import "database/sql"

type SQLite struct {
	// Add SQLite specific fields here

	DB *sql.DB
}
