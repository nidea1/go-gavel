package main

import (
	"context"
	"log/slog"
	"net"

	"github.com/nidea1/go-gavel/pkg/database"
	"github.com/nidea1/go-gavel/pkg/logger"
	pb "github.com/nidea1/go-gavel/proto/auth"
	"github.com/nidea1/go-gavel/services/auth/config"
	"github.com/nidea1/go-gavel/services/auth/internal/api/grpcservice"
	"github.com/nidea1/go-gavel/services/auth/internal/repository/postgres"
	"github.com/nidea1/go-gavel/services/auth/internal/usecase/command"
	"google.golang.org/grpc"
)

func main() {
	cfg := config.LoadConfig()
	logger.InitLogger(cfg.Server.Env)
	slog.Debug("Config loaded", "config", cfg)

	slog.Info("Starting auth service")
	slog.Info("Server is running on port", "port", cfg.Server.Port)

	pool, err := database.NewPgxPool(context.Background(), database.DBConfig{
		Host:              cfg.Database.Host,
		Port:              cfg.Database.Port,
		User:              cfg.Database.User,
		Password:          cfg.Database.Password,
		DBName:            cfg.Database.DBName,
		SSLMode:           cfg.Database.SSLMode,
		MinConns:          cfg.Database.MinConns,
		MaxConns:          cfg.Database.MaxConns,
		MaxConnLifetime:   cfg.Database.MaxConnLifetime,
		MaxConnIdleTime:   cfg.Database.MaxConnIdleTime,
		HealthCheckPeriod: cfg.Database.HealthCheckPeriod,
		ConnectionTimeout: cfg.Database.ConnectionTimeout,
	})
	if err != nil {
		slog.Error(err.Error())
	}
	defer pool.Close()

	slog.Info("Connected to database", "database", cfg.Database.DBName)

	userRepo := postgres.NewPostgresUserRepository(pool)

	// Will add kafka producer here

	registerUC := command.NewRegisterUseCase(userRepo)

	authService := grpcservice.NewAuthService(registerUC)

	grpcServer := grpc.NewServer()
	pb.RegisterAuthServiceServer(grpcServer, authService)

	listener, err := net.Listen("tcp", ":"+cfg.Server.Port)
	if err != nil {
		slog.Error(err.Error())
	}

	slog.Info("gRPC server is running")
	slog.Info("Listening on port", "port", cfg.Server.Port)
	if err := grpcServer.Serve(listener); err != nil {
		slog.Error(err.Error())
	}
}
