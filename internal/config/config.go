package config

import "github.com/joho/godotenv"

type GRPCConfig interface {
	Address() string
}

type PGConfig interface {
	DSN() string
}

type JWTConfig interface {
	Secret() string
	ExpiredRefresh() int
	ExpiredAccess() int
}

func Load(path string) error {
	err := godotenv.Load(path)
	if err != nil {
		return err
	}
	return nil
}
