package database

import (
	"github.com/ajaysinghnp/dcabinet/lib/config"
)

type Database interface {
	Connect(config *config.Config) error
	Close() error
	InitializeTables() error
}
