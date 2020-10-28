package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
	"github.com/woodchuckchoi/sweetpet/db"
	"github.com/woodchuckchoi/sweetpet/handler"
	"github.com/woodchuckchoi/sweetpet/route"
)

func main() {
	e := echo.New()
	e.Logger.SetLevel(log.ERROR)

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"}, // to modify later
		AllowMethods: []string{echo.GET, echo.HEAD, echo.PUT, echo.PATCH, echo.POST, echo.DELETE},
	}))

	db := db.Init()
	defer db.Close()

	h := &handler.Handler{DB: db}

	route.Match(e, h)

	e.Logger.Fatal(e.Start(":8080"))
}
