package database

import (
	"context"
	"log/slog"

	"fmt"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

type DBConfig struct {
	Host              string
	Port              int
	User              string
	Password          string
	DBName            string
	SSLMode           string
	MinConns          uint8
	MaxConns          uint8
	MaxConnLifetime   time.Duration
	MaxConnIdleTime   time.Duration
	HealthCheckPeriod time.Duration
	ConnectionTimeout time.Duration
}

func NewPgxPool(ctx context.Context, cfg DBConfig) (*pgxpool.Pool, error) {

	connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s",
		cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.DBName, cfg.SSLMode)

	config, err := pgxpool.ParseConfig(connStr)
	if err != nil {
		return nil, err
	}

	config.MinConns = cfg.MinConns
	config.MaxConns = cfg.MaxConns
	config.MaxConnLifetime = cfg.MaxConnLifetime * time.Hour
	config.MaxConnIdleTime = cfg.MaxConnIdleTime * time.Minute
	config.HealthCheckPeriod = cfg.HealthCheckPeriod * time.Minute
	config.ConnConfig.ConnectTimeout = cfg.ConnectionTimeout * time.Second

	config.AfterConnect = func() {
		slog.Info("Connected to database", "database", cfg.DBName)
	}

	config.BeforeClose = func(conn *pgxpool.Conn) {
		slog.Info("Closing connection to database", "database", cfg.DBName)
	}

	pool, err := pgxpool.NewWithConfig(ctx, config)
	if err != nil {
		return nil, err
	}

	return pool, nil
}
