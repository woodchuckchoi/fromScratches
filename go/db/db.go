package db

import (
	"fmt"
	"os"

	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

func Init() *sql.DB {
	// Init DB
	const (
		address string = "127.0.0.1:3306"
		dbUser  string = "root"
		dbName  string = "sweetpet"
	)

	var dbEndpoint string = fmt.Sprintf("%s:%s@tcp(%s)/%s?parseTime=true", dbUser, os.Getenv("DBPASS"), address, dbName)

	db, _ := sql.Open("mysql", dbEndpoint)

	db.SetMaxOpenConns(100)
	db.SetMaxIdleConns(100)

	return db
}
