package main

import (
	"log/slog"

	"github.com/EwanGreer/scaleable-e-commerce/internal/slogger"
)

func main() {
	slogger.InitGlobalSlogger(slog.LevelInfo)
	slog.Info("Users Called")
}
