package main

import (
	"log"
   "os"

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

	router.GET("/task/:taskId", api.GetTaskById)
	router.GET("/tasks/all", api.GetTasks)
	router.GET("/tasks/archived", api.GetArchivedTasks)
	router.GET("/tasks/count", api.CountAllTasks)
	router.GET("/tasks/late", api.GetLateTasks)
	router.GET("/tasks/late/count", api.CountAllLateTasks)

	router.DELETE("/task/:taskId", api.DeleteTask)

	router.PATCH("/task/done/:taskId", api.TaskDone)
	router.PATCH("/task/archive/:taskId", api.ArchiveTask)
	router.PATCH("/task/unarchive/:taskId", api.UnArchiveTask)

	router.POST("/create", api.CreateTask)

	router.PUT("/task/update", api.UpdateSpecificTask)

	ip := utils.GetLocalIP()
   router.Run(ip + ":" + os.Getenv("PORT"))
}
