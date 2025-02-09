package logger

import (
	"log/slog"
	"os"
)

var Log *slog.Logger

func InitLogger(env string) {
	switch env {
	case "production":
		Log = slog.New(slog.NewJSONHandler(os.Stdout, nil))
	default:
		Log = slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
			Level: slog.LevelDebug,
		}))
	}

	slog.SetDefault(Log)
}
