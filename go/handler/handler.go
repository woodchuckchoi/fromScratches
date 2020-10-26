package handler

import (
	"os"
	"database/sql"
)

type (
	Handler struct {
		DB *sql.DB
	}
)

var (
	secret string = os.Getenv("DBPASS")
)
