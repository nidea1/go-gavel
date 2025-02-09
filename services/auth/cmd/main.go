package main

import (
	"log/slog"

	"github.com/nidea1/go-gavel/pkg/logger"
	"github.com/nidea1/go-gavel/services/auth/config"
)

func main() {
	cfg := config.LoadConfig()
	logger.InitLogger(cfg.Server.Env)
	slog.Debug("Config loaded", "config", cfg)

	slog.Info("Starting auth service")
	slog.Info("Server is running on port", "port", cfg.Server.Port)
}
