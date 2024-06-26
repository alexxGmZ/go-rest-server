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

	router.GET("/task/all", api.GetTasks)
	router.GET("/task/:taskId", api.GetTaskById)
	router.GET("/task/late", api.GetLateTasks)
	router.POST("/create", api.CreateTask)
	router.DELETE("/task/:taskId", api.DeleteTask)

	ip := utils.GetLocalIP()
	router.Run(ip + ":8080")
}

