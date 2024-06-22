package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"go_rest/api"
	"go_rest/utils"
)

func main() {
	err := utils.ConnectDb()
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	router := gin.Default()
	router.GET("/tasks", api.GetTasks)
	router.Run("localhost:8080")
}
