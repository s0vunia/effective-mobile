package config

import (
	"time"

	"github.com/joho/godotenv"
)

// PGConfig - config for Postgres
type PGConfig interface {
	DSN() string
}

// HTTPConfig - config for HTTP server
type HTTPConfig interface {
	Address() string
	ReadHeaderTimeout() time.Duration
}

// StorageConfig - config for storage
type StorageConfig interface {
	Mode() string
}

// LoggerConfig - config for logger
type LoggerConfig interface {
	Level() string
	FileName() string
	MaxSize() int
	MaxAge() int
	MaxBackups() int
}

// Load - loads config from .env
func Load(path string) error {
	err := godotenv.Load(path)
	if err != nil {
		return err
	}

	return nil
}
