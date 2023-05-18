package test

import (
	"database/sql"
	"time"

	_ "github.com/lib/pq"
)

func SetupTestDB() *sql.DB {
	connStr := "user=postgres password=ilham dbname=belajar_golang_restful_api_test port=5432 host=localhost"
	db, err := sql.Open("postgres", connStr)

	if err != nil {
		panic(err)
	}

	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(20)
	db.SetConnMaxLifetime(60 * time.Minute)
	db.SetConnMaxIdleTime(10 * time.Minute)

	return db
}
