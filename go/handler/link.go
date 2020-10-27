package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/woodchuckchoi/sweetpet/model"
	"github.com/woodchuckchoi/sweetpet/util"
)

func (h *Handler) GenerateLink(c echo.Context) error {
	u := new(model.User)

	if err := c.Bind(u); err != nil {
		return err
	}

	newLink := &model.Link{UserID: u.ID}

	for {
		newLink.Link = util.RandomString(10)

		rows, err := h.DB.Query("SELECT id FROM link WHERE link = ?", newLink.Link)
		defer rows.Close()

		if err != nil {
			return err
		}

		if rows.Next() {
			continue
		}

		result, err := h.DB.Exec("INSERT INTO link(user_id, link) VALUES ( ?, ? )", newLink.UserID, newLink.Link)
		if err != nil {
			return err
		}

		return c.JSON(http.StatusOK, newLink)
	}
}

func (h *Handler) RetrieveLink(c echo.Context) error {
	l := new(model.Link)
	if err := c.Bind(l); err != nil {
		return err
	}

	row := h.DB.QueryRow("SELECT user_id FROM link WHERE link = ?", l.Link)
	if err = row.Scan(&l.UserID); err != nil {
		return err
	}

	row = h.DB.QueryRow("SELECT name, low, high FROM user WHERE id = ?", l.UserID)
	if err != nil {
		return err
	}
	u := &model.User{ID: l.UserID}
	if err = row.Scan(&u.Name, &u.Low, &u.High); err != nil {
		return err
	}

	rows, err := h.DB.Query("SELECT blood_sugar, ts FROM health WHERE user_id = ? ORDER BY ts ASC", u.UserID)
	if err != nil {
		return err
	}
	healthEntries := []model.Health{}
	for rows.Next() {
		var (
			bs int
			ts time.Time
		)

		if err = rows.Scan(&bs, &ts); err != nil {
			return err
		}

		healthEntries = append(healthEntries, model.Health{BloodSugar: bs, TS: model.JSONTime(ts)})
	}

	// Should I make a structure that contains both model.User and []model.Health to send back to front?
}
