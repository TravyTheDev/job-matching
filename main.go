package main

import (
	"fmt"
	"log"

	"github.com/TravyTheDev/job-matching/db"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	db, err := db.NewSqlStorage()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	fmt.Println("hello")
}
