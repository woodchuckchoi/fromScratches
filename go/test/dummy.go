package main

import (
	"fmt"
	"os"

	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/google/uuid"
	"github.com/woodchuckchoi/sweetpet/model"
)

func createUser(db *sql.DB) {
	names := []string{"Roy", "Hyuck", "Ddol"}
	low, high := 50, 120

	for i := 0; i < len(names); i++ {
		uid, err := uuid.NewUUID()
		u := &model.User{Name: names[i], UUID: uid.String(), Low: low, High: high}
		res, err := db.Exec("INSERT INTO user(name, uuid, low, high) VALUES( ?, ?, ?, ? )", u.Name, u.UUID, u.Low, u.High)
		fmt.Println(res, err)
	}
}

func selectUser(db *sql.DB) {
	rows, _ := db.Query("SELECT name, low, high FROM user")
	defer rows.Close()
	fmt.Println(rows.Next())
	for rows.Next() {
		var (
			n         string
			low, high int
		)
		_ = rows.Scan(&n, &low, &high)
		fmt.Println(n, low, high)
	}
}

func selectError(db *sql.DB) {
	rows, err := db.Query("SELECT name, low, high FROM user WHERE name = 'Yalu'")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(rows.Next())

	for rows.Next() {
		var (
			n         string
			low, high int
		)
		_ = rows.Scan(&n, &low, &high)
		fmt.Println(n, low, high)
	}
}

func main() {
	const (
		address string = "127.0.0.1:3306"
		dbUser  string = "root"
		dbName  string = "sweetpet"
	)

	var dbEndpoint string = fmt.Sprintf("%s:%s@tcp(%s)/%s", dbUser, os.Getenv("DBPASS"), address, dbName)

	// temporary test password hardwired,
	db, _ := sql.Open("mysql", dbEndpoint)
	defer db.Close()

	db.SetMaxOpenConns(100)
	db.SetMaxIdleConns(100)

	//createUser(db)
	selectUser(db)

	selectError(db)
}
