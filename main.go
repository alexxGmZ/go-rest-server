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

	router.GET("/tasks/all", api.GetTasks)
	router.GET("/task/:taskId", api.GetTaskById)
	router.GET("/tasks/late", api.GetLateTasks)

	router.DELETE("/task/:taskId", api.DeleteTask)

	router.PATCH("/task/done/:taskId", api.TaskDone)
	router.PATCH("/task/archive/:taskId", api.ArchiveTask)
	router.PATCH("/task/unarchive/:taskId", api.UnArchiveTask)

	router.POST("/create", api.CreateTask)

	ip := utils.GetLocalIP()
	router.Run(ip + ":8080")
}

