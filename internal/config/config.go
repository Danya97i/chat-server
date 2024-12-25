package config

import (
	"github.com/joho/godotenv"
)

// GRPCConfig - конфигурация для gRPC
type GRPCConfig interface {
	Address() string
}

// PGConfig - конфигурация для PostgreSQL
type PGConfig interface {
	DSN() string
}

// Load - загружает конфигурацию из файла .env
func Load(path string) error {
	return godotenv.Load(path)
}

// AuthClientConfig - конфигурация для сервиса авторизации
type AuthClientConfig interface {
	Address() string
	CertFile() string
}
