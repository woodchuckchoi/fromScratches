package main

import (
	"fmt"
	"os"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/woodchuckchoi/sweetpet/handler"
)

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"}, // to modify later
		AllowMethods: []string{echo.GET, echo.HEAD, echo.PUT, echo.PATCH, echo.POST, echo.DELETE},
	}))

	const (
		address string = "127.0.0.1:3306"
		dbUser string = "root"
		dbName string = "sweetpet"
	)

	var dbEndpoint string = fmt.Sprintf("%s:%s@tcp(%s)/%s", dbUser, os.Getenv("DBPASS"), address, dbName)

	// temporary test password hardwired, 
	db, _ := sql.Open("mysql", dbEndpoint)
	defer db.Close()

	db.SetMaxOpenConns(100)
	db.SetMaxIdleConns(100)

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello!")
	})

	e.Logger.Fatal(e.Start(":8080"))
}


// 1. How does go mod or dep work?
// 2. Directory structure for Go project?
