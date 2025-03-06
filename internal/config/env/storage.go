package env

import (
	"os"

	"github.com/pkg/errors"
	"github.com/s0vunia/effective-mobile/internal/config"
)

const storageModeEnvName = "STORAGE_MODE"

type storageConfig struct {
	mode string
}

func NewStorageConfig() (config.StorageConfig, error) {
	storageMode := os.Getenv(storageModeEnvName)
	if len(storageMode) == 0 {
		return nil, errors.New("storage mode not found")
	}

	return &storageConfig{
		mode: storageMode,
	}, nil
}

func (cfg *storageConfig) Mode() string {
	return cfg.mode
}
