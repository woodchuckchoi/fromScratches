package handler

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/woodchuckchoi/sweetpet/model"
	"github.com/woodchuckchoi/sweetpet/util"
)

/*
// RetrieveAll can use RetrieveOneDay with from 0 to timestamp MAX, need to refactor it
// link part and retrieve data part can be separated
func (h *Handler) RetrieveAllHealthEntries(c echo.Context) error {
	u := &model.User{Link: c.Param("link")}

	row := h.DB.QueryRow("SELECT id, low, high FROM user WHERE link = ?", u.Link)
	if err := row.Scan(&u.ID, &u.Low, &u.High); err != nil {
		return err
	}

	rows, err := h.DB.Query("SELECT blood_sugar, ts FROM health WHERE user_id = ? ORDER BY ts ASC", u.ID)
	if err != nil {
		return err
	}

	healthEntries := []model.Health{}
	for rows.Next() {
		entry := model.Health{}

		if err = rows.Scan(&entry.BloodSugar, &entry.Ts); err != nil {
			return err
		}

		healthEntries = append(healthEntries, entry)
	}

	type Res struct {
		User    model.User     `json:"user"`
		Entries []model.Health `json:"entries"`
	}

	return c.JSON(http.StatusOK, Res{User: *u, Entries: healthEntries})
}
*/

func (h *Handler) RetrieveRangedHealthEntries(c echo.Context) error {
	var err error

	var (
		from string = c.QueryParam("from")
		to   string = c.QueryParam("to")
	)

	u := new(model.User)
	row := h.DB.QueryRow("SELECT id, name, low, high FROM user WHERE link = ?", c.Param("link"))
	if err := row.Scan(&u.ID, &u.Name, &u.Low, &u.High); err != nil {
		return err
	}

	util.ToSqlTimeStamp(&from, &to)

	var query string
	if from != "" && to != "" {
		query = fmt.Sprintf("SELECT blood_sugar, ts FROM health WHERE user_id = %v AND ts >= %v AND ts <= %v ORDER BY ts ASC", u.ID, from, to)
	} else if from != "" {
		query = fmt.Sprintf("SELECT blood_sugar, ts FROM health WHERE user_id = %v AND ts >= %v ORDER BY ts ASC", u.ID, from)
	} else if to != "" {
		query = fmt.Sprintf("SELECT blood_sugar, ts FROM health WHERE user_id = %v AND ts <= %v ORDER BY ts ASC", u.ID, to)
	} else {
		query = fmt.Sprintf("SELECT blood_sugar, ts FROM health WHERE user_id = %v ORDER BY ts ASC", u.ID)
		fmt.Println("Right place")
	}

	rows, err := h.DB.Query(query)
	if err != nil {
		return err
	}

	healthEntries := []model.Health{}
	for rows.Next() {
		entry := model.Health{}
		if err = rows.Scan(&entry.BloodSugar, &entry.Ts); err != nil {
			return err
		}
		healthEntries = append(healthEntries, entry)
	}

	type Res struct {
		User    model.User     `json:"user"`
		Entries []model.Health `json:"entries"`
	}

	return c.JSON(http.StatusOK, Res{User: *u, Entries: healthEntries})
}

func (h *Handler) CommitEntry(c echo.Context) error {
	link := c.Param("link")

	health := new(model.Health)

	if err := c.Bind(health); err != nil {
		return err
	}

	_, err := h.DB.Exec("INSERT INTO health(user_id, blood_sugar, ts) VALUES( (SELECT id FROM user WHERE link = ?), ?, ? ) ", link, health.BloodSugar, health.Ts.ToString())
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, "success")
}

func (h *Handler) DeleteEntry(c echo.Context) error {
	link := c.Param("link")
	ts := c.Param("ts")

	util.ToSqlTimeStamp(&ts)

	query := fmt.Sprintf("DELETE FROM health WHERE user_id = (SELECT id FROM user WHERE link = '%v') AND ts = %v", link, ts)
	res, err := h.DB.Exec(query)
	if err != nil {
		return err
	}

	if val, _ := res.RowsAffected(); val > 0 {
		return c.JSON(http.StatusOK, "success")
	}

	return c.JSON(http.StatusBadRequest, "invalid request")
}

func (h *Handler) ModifyEntry(c echo.Context) error {
	link := c.Param("link")
	ts := c.Param("ts")

	util.ToSqlTimeStamp(&ts)

	health := new(model.Health)
	if err := c.Bind(health); err != nil {
		return err
	}

	query := fmt.Sprintf("UPDATE health SET blood_sugar = %v WHERE user_id = (SELECT id FROM user WHERE link = '%v') AND ts = %v", health.BloodSugar, link, ts)
	res, err := h.DB.Exec(query)
	if err != nil {
		return err
	}

	if val, _ := res.RowsAffected(); val > 0 {
		return c.JSON(http.StatusOK, "success")
	}

	return c.JSON(http.StatusBadRequest, "invalid request")
}
