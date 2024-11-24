package main

import (
	"log/slog"

	"github.com/EwanGreer/scaleable-e-commerce/internal/queues/kafka"
	"github.com/EwanGreer/scaleable-e-commerce/internal/slogger"
	"github.com/EwanGreer/scaleable-e-commerce/services/user/api"
	"github.com/EwanGreer/scaleable-e-commerce/services/user/config"
	"github.com/EwanGreer/scaleable-e-commerce/services/user/service"
)

func main() {
	slogger.InitGlobalSlogger(slog.LevelInfo)

	cfg := config.New()

	// Taking the first index of topics for now. Not sure if a record should be produced to multiple topics
	producer := kafka.NewProducer(cfg.Kafka.Producer.Topics[0], cfg.Kafka.Producer.Brokers)

	handler := api.New()

	svc := service.NewService(cfg, producer, handler)
	svc.Start()
}
