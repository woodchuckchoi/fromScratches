package main

import (
	"github.com/kataras/iris/v12"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	app := iris.New()

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
