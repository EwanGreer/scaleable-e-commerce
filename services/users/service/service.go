package service

import (
	"log"
	"sync"

	"github.com/EwanGreer/scaleable-e-commerce/internal/queues/kafka"
	"github.com/EwanGreer/scaleable-e-commerce/services/users/config"
)

type UserService struct {
	Producer    kafka.Producer
	ServiceName string
}

func New(cfg *config.AppConfig, producer kafka.Producer) *UserService {
	return &UserService{
		ServiceName: cfg.ServiceName,
		Producer:    producer,
	}
}

func (s *UserService) Start() {
	log.Println("start called")

	var wg sync.WaitGroup

	wg.Add(1)

	go func() {
		s.Producer.Produce([]byte(`{"msg":"Hello Kafka"}`))
	}()

	wg.Wait()
}
