package main

import (
	"fmt"
	"os"

	"github.com/kataras/iris/v12"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	app := iris.New()

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

	booksAPI := app.Party("/books")
	{
		booksAPI.Use(iris.Compression)

		booksAPI.Get("/", list)
		booksAPI.Post("/", create)
	}

	app.Listen(":8080")
}

type Book struct {
	Title string `json:"title"`
}

func list(ctx iris.Context) {
	books := []Book{
		{"Mastering Concurrency in Go"},
		{"Go Design Patterns"},
		{"Black Hat Go"},
	}

	ctx.JSON(books)
}

func create(ctx iris.Context) {
	var b Book
	err := ctx.ReadJSON(&b)

	if err != nil {
		ctx.StopWithProblem(iris.StatusBadRequest, iris.NewProblem().Title("Book creation failure").DetailErr(err))
		return
	}

	println("Received Book: " + b.Title)

	ctx.StatusCode(iris.StatusCreated)
}
