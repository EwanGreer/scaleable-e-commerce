package main

import (
	"log/slog"
	"os"

	"github.com/EwanGreer/scaleable-e-commerce/internal/queues/kafka"
	"github.com/EwanGreer/scaleable-e-commerce/services/user/api"
	"github.com/EwanGreer/scaleable-e-commerce/services/user/config"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type UserService struct {
	handler     *api.Handler
	Producer    kafka.Producer
	ServiceName string
	ListenAddr  string
	Port        string
}

func NewService(cfg *config.AppConfig, producer kafka.Producer, h *api.Handler) *UserService {
	return &UserService{
		ServiceName: cfg.ServiceName,
		Producer:    producer,
		ListenAddr:  cfg.Server.ListenAddr,
		Port:        cfg.Server.Port,
		handler:     h,
	}
}

func (s *UserService) Start() {
	e := echo.New()
	e.HideBanner = true

	MountRoutes(e, s.handler)

	if err := e.Start(s.ListenAddr + s.Port); err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}
}

func MountRoutes(e *echo.Echo, h *api.Handler) {
	e.Use(middleware.CORS())
	e.Use(middleware.RequestID())
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	// TODO: JWT auth

	api := e.Group("/api")
	v1 := api.Group("/v1")

	v1.GET("/health", func(e echo.Context) error {
		return e.JSON(200, echo.Map{"healthy": true})
	})

	v1.POST("/user", h.CreateUser)
	v1.GET("/user/:userId", h.GetUserById)
	v1.PUT("/user", h.UpdateUser)
	v1.DELETE("/user/:userId", h.DeleteUser)
}
