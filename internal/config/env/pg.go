package env

import (
	"errors"
	"os"

	"github.com/s0vunia/effective-mobile/internal/config"
)

const (
	dsnEnvName = "PG_DSN"
)

type pgConfig struct {
	dsn string
}

func NewPGConfig() (config.PGConfig, error) {
	dsn := os.Getenv(dsnEnvName)
	if len(dsn) == 0 {
		return nil, errors.New("pg dsn not found")
	}

	return &pgConfig{
		dsn: dsn,
	}, nil
}

func (cfg *pgConfig) DSN() string {
	return cfg.dsn
}
