package config

import (
	"github.com/nidea1/go-gavel/pkg/utils"
)

type ServerConfig struct {
	Port string
	Env  string
}

type DatabaseConfig struct {
	Host              string
	Port              int
	User              string
	Password          string
	DBName            string
	SSLMode           string
	MinConns          int
	MaxConns          int
	MaxConnLifetime   int
	MaxConnIdleTime   int
	HealthCheckPeriod int
	ConnectionTimeout int
}

type JWTConfig struct {
	SecretKey       string
	AccessTokenTTL  int
	RefreshTokenTTL int
}

type Config struct {
	Server   ServerConfig
	Database DatabaseConfig
	JWT      JWTConfig
	// Kafka will be added here
}

func LoadConfig() *Config {
	return &Config{
		Server: ServerConfig{
			Port: utils.GetEnvString("SERVER_PORT", "8080"),
			Env:  utils.GetEnvString("ENV", "development"),
		},
		Database: DatabaseConfig{
			Host:              utils.GetEnvString("POSTGRES_HOST", "localhost"),
			Port:              utils.GetEnvInt("POSTGRES_PORT", 5432),
			User:              utils.GetEnvString("POSTGRES_USER", "postgres"),
			Password:          utils.GetEnvString("POSTGRES_PASSWORD", "password"),
			DBName:            utils.GetEnvString("POSTGRES_DB", "auth"),
			SSLMode:           utils.GetEnvString("POSTGRES_SSL_MODE", "disable"),
			MinConns:          utils.GetEnvInt("POSTGRES_MIN_CONNS", 0),
			MaxConns:          utils.GetEnvInt("POSTGRES_MAX_CONNS", 10),
			MaxConnLifetime:   utils.GetEnvInt("POSTGRES_MAX_CONN_LIFETIME", 30),
			MaxConnIdleTime:   utils.GetEnvInt("POSTGRES_MAX_CONN_IDLE_TIME", 10),
			HealthCheckPeriod: utils.GetEnvInt("POSTGRES_HEALTH_CHECK_PERIOD", 60),
			ConnectionTimeout: utils.GetEnvInt("POSTGRES_CONNECTION_TIMEOUT", 5),
		},
		JWT: JWTConfig{
			SecretKey:       utils.GetEnvString("JWT_SECRET_KEY", "secret"),
			AccessTokenTTL:  utils.GetEnvInt("JWT_ACCESS_TOKEN_TTL", 15),
			RefreshTokenTTL: utils.GetEnvInt("JWT_REFRESH_TOKEN_TTL", 30),
		},
	}
}
