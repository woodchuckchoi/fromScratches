package handler

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/woodchuckchoi/sweetpet/model"
	"github.com/woodchuckchoi/sweetpet/util"
)

func (h *Handler) Register(c echo.Context) error {
	u := new(model.User)

	if err := c.Bind(u); err != nil {
		return &echo.HTTPError{Code: http.StatusBadRequest, Message: "inappropriate request"}
	}

	uid, err := uuid.NewUUID()
	u.UUID = uid.String()

	if u.Name == "" || len(u.Name) > 20 {
		return &echo.HTTPError{Code: http.StatusBadRequest, Message: "invalid name"}
	}

	_, err = h.DB.Exec("INSERT INTO user(name, uuid, low, high) VALUES( ?, ?, ?, ? )", u.Name, u.UUID, u.Low, u.High)

	if err != nil {
		return &echo.HTTPError{Code: http.StatusBadRequest, Message: err}
	}

	return c.JSON(http.StatusCreated, u)
}

func (h *Handler) ModifyThreshold(c echo.Context) error {
	u := new(model.User)
	if err := c.Bind(u); err != nil {
		return err
	}

	_, err := h.DB.Exec("UPDATE user SET low = ?, high = ? WHERE id = ?", u.Low, u.High, u.ID)

	if err != nil {
		return &echo.HTTPError{Code: http.StatusBadRequest, Message: err}
	}

	return c.JSON(http.StatusOK, u)
}

func (h *Handler) GenerateLink(c echo.Context) error {
	u := new(model.User)

	if err := c.Bind(u); err != nil {
		return err
	}

	for {
		u.Link = util.RandomString(10)

		// Below Transaction should be isolated in production
		// From HERE
		rows, err := h.DB.Query("SELECT id FROM user WHERE link = ?", u.Link) // Since index(primary key) is stored with all column, it is okay to select id, not link.
		if err != nil {
			return err
		}

		if rows.Next() {
			continue
		}

		_, err = h.DB.Exec("UPDATE user SET link = ? WHERE id = ?", u.Link, u.ID)
		if err != nil {
			return err
		}
		// To HERE

		return c.JSON(http.StatusCreated, u)
	}
}
