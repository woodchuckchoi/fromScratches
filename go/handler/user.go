package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/woodchuckchoi/sweetpet/model"
	"github.com/google/uuid"
)

func (h *Handler) Register(c echo.Context) error {
	u := &model.User{UUID: uuid.NewUUID().String()}
	if err = c.Bind(u); err != nil {
		return echo.HTTPError{Code: http.StatusBadRequest, Message: "inappropriate request"}
	}

	if u.Name == "" || len(u.Name) > 20 {
		return echo.HTTPError{Code: http.StatusBadRequest, Message: "invalid name"}
	}

	_, err := h.DB.Exec("INSERT INTO user(name, uuid, low, high) VALUES( ?, ?, ?, ? )", u.Name, u.UUID, u.Low, u.High)

	if err != nil {
		return echo.HTTPError{Code: http.StatusBadRequest, Message: err.String()}
	}

	return c.JSON(http.StatusOK, u)
}

