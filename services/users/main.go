package main

import (
	"log/slog"

	"github.com/EwanGreer/scaleable-e-commerce/internal/queues/kafka"
	"github.com/EwanGreer/scaleable-e-commerce/internal/slogger"
	"github.com/EwanGreer/scaleable-e-commerce/services/users/api"
	"github.com/EwanGreer/scaleable-e-commerce/services/users/config"
)

func main() {
	slogger.InitGlobalSlogger(slog.LevelInfo)

	cfg := config.New()

	// Taking the first index of topics for now. Not sure if a record should be produced to multiple topics
	producer := kafka.NewProducer(cfg.Kafka.Producer.Topics[0], cfg.Kafka.Producer.Brokers)

	handler := api.New()

	svc := NewService(cfg, producer, handler)
	svc.Start()
}
