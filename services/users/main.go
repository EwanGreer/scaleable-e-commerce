package main

import (
	"log/slog"

	"github.com/EwanGreer/scaleable-e-commerce/internal/slogger"
	"github.com/EwanGreer/scaleable-e-commerce/services/users/config"
	"github.com/EwanGreer/scaleable-e-commerce/services/users/service"
)

func main() {
	slogger.InitGlobalSlogger(slog.LevelInfo)

	cfg := config.New("development")

	slog.Info("Init Users")

	svc := service.New(cfg, "tmp")
	svc.Start()
}
