package controller

import (
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	myerr "github.com/woodchuckchoi/fromScratches/src/core/err"
	"github.com/woodchuckchoi/fromScratches/src/core/model"
	"github.com/woodchuckchoi/fromScratches/src/core/service"
)

func requestAnalytics(c echo.Context) error {
	list := []model.Crop{}
	resp := []model.Harvest{}

	if err := c.Bind(list); err != nil {
		log.Fatal(err)
		return myerr.InvalidParameterError
	}

	resp = service.RequestAnalytics(list)

	return c.JSON(http.StatusOK, resp)
}
