package main

import (
	"log/slog"

	"github.com/EwanGreer/scaleable-e-commerce/internal/slogger"
	"github.com/EwanGreer/scaleable-e-commerce/services/users/config"
)

func main() {
	slogger.InitGlobalSlogger(slog.LevelInfo)

	cfg := config.New("development")
	_ = cfg

	slog.Info("Init Users")
}
