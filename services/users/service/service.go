package service

import (
	"log/slog"

	"github.com/EwanGreer/scaleable-e-commerce/services/users/config"
)

type UserService struct {
	RecordProcessor any
	ServiceName     string
}

func New(cfg *config.AppConfig, processor any) *UserService {
	return &UserService{
		ServiceName:     cfg.ServiceName,
		RecordProcessor: processor,
	}
}

func (*UserService) Start() {
	slog.Info("Started Consumer Loop")
}
