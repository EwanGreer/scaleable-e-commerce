package api

import (
	"errors"
	"strconv"

	"github.com/EwanGreer/scaleable-e-commerce/services/user/repo"
	"github.com/labstack/echo/v4"
)

type Handler struct {
	DB *repo.Queries
}

func NewHandler(queries *repo.Queries) *Handler {
	return &Handler{DB: queries}
}

func (h *Handler) GetUserById(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("userId"))
	if err != nil {
		return err
	}

	if id < 0 {
		return errors.New("input: id is less than 0")
	}

	user, err := h.DB.GetUserByID(c.Request().Context(), int32(id))
	if err != nil {
		return err
	}

	return c.JSON(200, user)
}

func (h *Handler) CreateUser(c echo.Context) error { return nil }
func (h *Handler) UpdateUser(c echo.Context) error { return nil }
func (h *Handler) DeleteUser(c echo.Context) error { return nil }
