package slogger

import (
	"log/slog"
	"os"
)

func InitGlobalSlogger(level slog.Level) {
	logger := slog.New(slog.NewJSONHandler(os.Stderr, &slog.HandlerOptions{
		Level: level,
	}))
	slog.SetDefault(logger)
}
