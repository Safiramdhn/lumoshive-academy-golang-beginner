package database

import (
	"database/sql"
	_ "github.com/lib/pq"
)

func InitDb() (*sql.DB, error) {
	// Open the database
	connStr := "host=localhost user=postgres password=postgres dbname=ojek-online sslmode=disable"
	db, err := sql.Open("postgres", connStr)

	return db, err
}
