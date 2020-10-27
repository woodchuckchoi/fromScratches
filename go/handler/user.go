package handler

import (
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/woodchuckchoi/sweetpet/model"
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

func (h *Handler) ModifyThreshold(c echo.Context) error {
	u := new(model.User)
	if err = c.Bind(u); err != nil {
		return err
	}

	_, err := h.DB.Exec("UPDATE user(low, high) VALUES( ?, ?) WHERE id = ? AND uuid = ?", u.Low, u.High, u.Id, u.UUID)

	if err != nil {
		return echo.HTTPError{Code: http.StatusBadRequest, Message: err.String()}
	}

	return c.JSON(http.StatusOK, u)
}
