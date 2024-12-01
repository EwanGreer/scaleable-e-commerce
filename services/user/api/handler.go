package api

import (
	"errors"
	"strconv"

	"github.com/EwanGreer/scaleable-e-commerce/services/user/repo"
	"github.com/labstack/echo/v4"
)

type Handler struct {
	DB repo.Querier
}

func NewHandler(queries repo.Querier) *Handler {
	return &Handler{DB: queries}
}

func (h *Handler) Health(c echo.Context) error {
	// TODO: call out to DB ping/health
	return c.JSON(200, map[string]any{"healthy": true})
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

func (h *Handler) CreateUser(c echo.Context) error {
	var user repo.CreateUserParams // NOTE: is there an easy way to validate this with struct tags?

	if err := c.Bind(&user); err != nil {
		return err
	}

	u, err := h.DB.CreateUser(c.Request().Context(), user)
	if err != nil {
		return err
	}

	return c.JSON(200, map[string]any{"id": u.ID})
}

func (h *Handler) UpdateUser(c echo.Context) error { return nil }
func (h *Handler) DeleteUser(c echo.Context) error { return nil }
