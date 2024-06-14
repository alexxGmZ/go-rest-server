package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	// "github.com/lib/pq"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file")
	}

	postgresConnURL := os.Getenv("POSTGRESURL")

	log.Println(postgresConnURL)
}
