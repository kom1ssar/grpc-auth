package config

import (
	"github.com/joho/godotenv"
	"time"
)

type GRPCConfig interface {
	Address() string
}

type PGConfig interface {
	DSN() string
}

type JWTConfig interface {
	AccessSecret() string
	RefreshSecret() string
	SecondsExpiredRefresh() time.Duration
	SecondsExpiredAccess() time.Duration
}

func Load(path string) error {
	err := godotenv.Load(path)
	if err != nil {
		return err
	}
	return nil
}
