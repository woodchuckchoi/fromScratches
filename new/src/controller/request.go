package controller

import (
	"log"

	"github.com/labstack/echo/v4"
	myerr "github.com/woodchuckchoi/fromScratches/src/core/err"
	"github.com/woodchuckchoi/fromScratches/src/core/model"
)

func requestAnalytics(c echo.Context) error {
	list := []model.Crop{}

	if err := c.Bind(list); err != nil {
		log.Fatal(err)
		return myerr.InvalidParameterError
	}
	var resReceiver = make(<-chan model.Harvest, len(list))

	for _, crop := range list {
		go func(c model.Crop, receiver <-chan model.Harvest) {

		}(crop, resReceiver)
	}

	select {
	case h := <-resReceiver:

	}
	return nil
}
