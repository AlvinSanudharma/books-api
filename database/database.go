package database

import (
	"database/sql"

	_ "github.com/lib/pq"
)

func InitDb() {
	dsn := "postgres://postgres@localhost:5432/postgres?sslmode=disable"
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}
}
