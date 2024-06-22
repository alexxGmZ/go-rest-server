package utils

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var (
	DB *sql.DB
)

func ConnectDb() error {
	err := godotenv.Load()
	if err != nil {
		return fmt.Errorf("error loading .env file: %w", err)
	}

	postgresConnURL := os.Getenv("POSTGRESURL")

	DB, err = sql.Open("postgres", postgresConnURL)
	if err != nil {
		return fmt.Errorf("error connecting to database: %w", err)
	}

	if err := DB.Ping(); err != nil {
		return fmt.Errorf("error pinging database: %w", err)
	}

	log.Println("Connected to database")
	return nil
}
