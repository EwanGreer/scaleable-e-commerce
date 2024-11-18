package api

import "github.com/labstack/echo/v4"

type Handler struct {
	DB any // Likely a postgres DB
}

func New() *Handler {
	return &Handler{}
}

func (h Handler) GetUserById(c echo.Context) error { return nil }
func (h Handler) CreateUser(c echo.Context) error  { return nil }
func (h Handler) UpdateUser(c echo.Context) error  { return nil }
func (h Handler) DeleteUser(c echo.Context) error  { return nil }
