package main

import (
	"context"
	"log/slog"

	"github.com/EwanGreer/scaleable-e-commerce/internal/queues/kafka"
	"github.com/EwanGreer/scaleable-e-commerce/internal/slogger"
	"github.com/EwanGreer/scaleable-e-commerce/services/user/api"
	"github.com/EwanGreer/scaleable-e-commerce/services/user/config"
	"github.com/EwanGreer/scaleable-e-commerce/services/user/repo"
	"github.com/EwanGreer/scaleable-e-commerce/services/user/service"
	"github.com/jackc/pgx/v5"
)

func main() {
	slogger.InitGlobalSlogger(slog.LevelInfo)

	cfg := config.Load()

	// Taking the first index of topics for now. Not sure if a record should be produced to multiple topics
	producer := kafka.NewProducer(cfg.Kafka.Producer.Topics[0], cfg.Kafka.Producer.Brokers)

	ctx := context.Background()

	conn, err := pgx.Connect(ctx, cfg.Database.ConnectionString)
	if err != nil {
		panic(err)
	}
	defer conn.Close(ctx)

	db := repo.New(conn)

	handler := api.NewHandler(db)

	svc := service.NewService(cfg, producer, handler)
	svc.Start()
}
