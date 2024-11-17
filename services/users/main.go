package main

import (
	"log/slog"

	"github.com/EwanGreer/scaleable-e-commerce/internal/queues/kafka"
	"github.com/EwanGreer/scaleable-e-commerce/internal/slogger"
	"github.com/EwanGreer/scaleable-e-commerce/services/users/config"
	"github.com/EwanGreer/scaleable-e-commerce/services/users/service"
)

func main() {
	slogger.InitGlobalSlogger(slog.LevelInfo)

	cfg := config.New("development")

	slog.Info("Init Users")

	// taking the first index of topics for now. Not sure if a record should be produced to multiple topics
	producer := kafka.NewProducer(cfg.Kafka.Producer.Topics[0], cfg.Kafka.Producer.Brokers)

	svc := service.New(cfg, producer)
	svc.Start()
}
