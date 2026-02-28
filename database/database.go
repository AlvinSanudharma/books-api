package database

import (
	"database/sql"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func InitDb() {
	dsn := "postgres://postgres@localhost:5432/books-api?sslmode=disable"
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	DB = db
}
