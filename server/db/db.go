package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"
)

func NewSqlStorage() (*sql.DB, error) {
	DB_HOST := os.Getenv("DB_HOST")
	DB_USER := os.Getenv("DB_USER")
	DB_NAME := os.Getenv("DB_NAME")
	DB_PASSWORD := os.Getenv("DB_PASSWORD")
	DB_SSLMODE := os.Getenv("DB_SSLMODE")

	uri := fmt.Sprintf("postgres://%s:%s@%s:5432/%s?sslmode=%s",
		DB_USER,
		DB_PASSWORD,
		DB_HOST,
		DB_NAME,
		DB_SSLMODE,
	)

	db, err := sql.Open("postgres", uri)
	if err != nil {
		log.Fatal(err)
	}
	return db, nil
}
