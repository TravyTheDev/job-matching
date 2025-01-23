package main

import (
	"log"
	"os"

	"github.com/TravyTheDev/job-matching/api"
	"github.com/TravyTheDev/job-matching/db"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	port := os.Getenv("PORT")
	db, err := db.NewSqlStorage()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	server := api.NewAPIServer(port, db)

	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}
