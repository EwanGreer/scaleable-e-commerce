package service

import (
	"fmt"
	"log/slog"
	"os"

	"github.com/EwanGreer/scaleable-e-commerce/internal/queues/kafka"
	"github.com/EwanGreer/scaleable-e-commerce/services/users/config"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type UserService struct {
	Producer    kafka.Producer
	ServiceName string
	ListenAddr  string
	Port        string
}

func New(cfg *config.AppConfig, producer kafka.Producer) *UserService {
	return &UserService{
		ServiceName: cfg.ServiceName,
		Producer:    producer,
		ListenAddr:  cfg.Server.ListenAddr,
		Port:        cfg.Server.Port,
	}
}

func (s *UserService) Start() {
	e := echo.New()
	e.HideBanner = true

	MountRoutes(e)

	if err := e.Start(fmt.Sprintf("%s%s", s.ListenAddr, s.Port)); err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}
}

func MountRoutes(e *echo.Echo) {
	e.Use(middleware.CORS())
	e.Use(middleware.RequestID())
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	// TODO: jwt auth

	api := e.Group("/api")
	v1 := api.Group("/v1")

	v1.GET("/health", func(e echo.Context) error {
		return e.JSON(200, echo.Map{"healthy": true})
	})
}
