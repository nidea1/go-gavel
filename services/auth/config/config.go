package config

import (
	"github.com/nidea1/go-gavel/pkg/utils"
)

type Config struct {
	Server struct {
		Port string
	}
	Database struct {
		Host     string
		Port     int
		User     string
		Password string
		DBName   string
		SSLMode  string
	}
	JWT struct {
		SecretKey       string
		AccessTokenTTL  int
		RefreshTokenTTL int
	}
	// Kafka will be added here
}

func LoadConfig() *Config {
	return &Config{
		Server: struct {
			Port string
		}{
			Port: utils.GetEnvString("SERVER_PORT", "8080"),
		},
		Database: struct {
			Host     string
			Port     int
			User     string
			Password string
			DBName   string
			SSLMode  string
		}{
			Host:     utils.GetEnvString("POSTGRES_HOST", "localhost"),
			Port:     utils.GetEnvInt("POSTGRES_PORT", 5432),
			User:     utils.GetEnvString("POSTGRES_USER", "postgres"),
			Password: utils.GetEnvString("POSTGRES_PASSWORD", "password"),
			DBName:   utils.GetEnvString("POSTGRES_DB", "auth"),
			SSLMode:  utils.GetEnvString("POSTGRES_SSL_MODE", "disable"),
		},
		JWT: struct {
			SecretKey       string
			AccessTokenTTL  int
			RefreshTokenTTL int
		}{
			SecretKey:       utils.GetEnvString("JWT_SECRET_KEY", "secret"),
			AccessTokenTTL:  utils.GetEnvInt("JWT_ACCESS_TOKEN_TTL", 15),
			RefreshTokenTTL: utils.GetEnvInt("JWT_REFRESH_TOKEN_TTL", 30),
		},
	}
}
