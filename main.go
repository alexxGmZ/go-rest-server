package main

import (
	"log"
	"net"

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
	router.GET("/tasks/:taskId", api.GetTaskById)

	ip := getLocalIP()
	router.Run(ip + ":8080")
}

func getLocalIP() string {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		log.Println(err)
		return "localhost"
	}

	for _, addr := range addrs {
		if ipnet, ok := addr.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return ipnet.IP.String()
			}
		}
	}

	return "localhost"
}
