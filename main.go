package main

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"

	// "github.com/lib/pq"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}

	fmt.Println(os.Getenv("TEST"))
}
