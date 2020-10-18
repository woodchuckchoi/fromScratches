package main

import (
	"fmt"
	"os"
	"net/http"

	"github.com/labstack/echo/v4"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	e := echo.New()

	const (
		address string = "127.0.0.1:3306"
		dbUser string = "root"
		dbName string = "sweetpet"
	)

	var dbEndpoint string = fmt.Sprintf("%s:%s@tcp(%s)/%s", dbUser, os.Getenv("DBPASS"), address, dbName)

	// temporary test password hardwired, 
	db, _ := sql.Open("mysql", dbEndpoint)
	defer db.Close()

	var version string
	db.QueryRow("SELECT VERSION()").Scan(&version)
	println("Connected to:", version)

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello!")
	})

	e.Logger.Fatal(e.Start(":8080"))
}

type Book struct {
	Title string `json:"title"`
}

