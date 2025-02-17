package config

import (
	"time"

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
	MinConns          int32
	MaxConns          int32
	MaxConnLifetime   time.Duration
	MaxConnIdleTime   time.Duration
	HealthCheckPeriod time.Duration
	ConnectionTimeout time.Duration
}

type JWTConfig struct {
	SecretKey       string
	AccessTokenTTL  time.Duration
	RefreshTokenTTL time.Duration
}

type RedisConfig struct {
	Addr     string
	Password string
	DB       uint8
}

type Config struct {
	Server   ServerConfig
	Database DatabaseConfig
	JWT      JWTConfig
	Redis    RedisConfig
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
			MinConns:          int32(utils.GetEnvInt("POSTGRES_MIN_CONNS", 0)),
			MaxConns:          int32(utils.GetEnvInt("POSTGRES_MAX_CONNS", 10)),
			MaxConnLifetime:   time.Duration(utils.GetEnvInt("POSTGRES_MAX_CONN_LIFETIME", 30)) * time.Minute,
			MaxConnIdleTime:   time.Duration(utils.GetEnvInt("POSTGRES_MAX_CONN_IDLE_TIME", 10)) * time.Minute,
			HealthCheckPeriod: time.Duration(utils.GetEnvInt("POSTGRES_HEALTH_CHECK_PERIOD", 60)) * time.Minute,
			ConnectionTimeout: time.Duration(utils.GetEnvInt("POSTGRES_CONNECTION_TIMEOUT", 5)) * time.Second,
		},
		JWT: JWTConfig{
			SecretKey:       utils.GetEnvString("JWT_SECRET_KEY", "secret"),
			AccessTokenTTL:  time.Duration(utils.GetEnvInt("JWT_ACCESS_TOKEN_TTL", 15)) * time.Minute,
			RefreshTokenTTL: time.Duration(utils.GetEnvInt("JWT_REFRESH_TOKEN_TTL", 24*60)) * time.Minute,
		},
		Redis: RedisConfig{
			Addr:     utils.GetEnvString("REDIS_ADDR", "localhost:6379"),
			Password: utils.GetEnvString("REDIS_PASSWORD", ""),
			DB:       uint8(utils.GetEnvInt("REDIS_DB", 0)),
		},
	}
}
