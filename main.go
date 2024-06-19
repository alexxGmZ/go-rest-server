package main

import (
	"fmt"
	"log"
	"os"

	"database/sql"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"net/http"
)

var (
	postgresConnURL string
	db              *sql.DB
)

type Task struct {
	TaskID      int    `json:"task_id"`
	Description string `json:"description"`
	Status      string `json:"status"`
	Deadline    string `json:"deadline"`
	DateAdded   string `json:"date_added"`
}

func main() {
	err := connectDb()
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	router := gin.Default()
	router.GET("/tasks", getToDos)
	router.Run("localhost:8080")
}

func getToDos(c *gin.Context) {
	query := `
		SELECT task_id, description, status, deadline, date_added
		FROM Tasks
	`

	rows, err := db.Query(query)
	if err != nil {
		c.JSON(
			http.StatusInternalServerError,
			gin.H{"error": "Failed to query tasks"},
		)
		return
	}

	var tasks []Task
	for rows.Next() {
		var task Task
		if err := rows.Scan(&task.TaskID, &task.Description, &task.Status, &task.Deadline, &task.DateAdded); err != nil {
			c.JSON(
				http.StatusInternalServerError,
				gin.H{"error": "Failed to scan task"},
			)
			return
		}
		fmt.Println(task)
		tasks = append(tasks, task)
	}

	if err := rows.Err(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error iterating through tasks"})
		return
	}

	c.JSON(http.StatusOK, tasks)
}

func connectDb() error {
	err := godotenv.Load()
	if err != nil {
		return fmt.Errorf("error loading .env file: %w", err)
	}

	postgresConnURL := os.Getenv("POSTGRESURL")

	db, err = sql.Open("postgres", postgresConnURL)
	if err != nil {
		return fmt.Errorf("error connecting to database: %w", err)
	}

	if err := db.Ping(); err != nil {
		return fmt.Errorf("error pinging database: %w", err)
	}

	log.Println("Connected to database:", postgresConnURL)
	return nil
}
