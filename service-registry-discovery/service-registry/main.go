package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type Service struct {
	Name string `json:"name"`
	Host string `json:"host"`
	Port int    `json:"port"`
}

var services []Service

func main() {
	router := gin.Default()
	router.POST("/register", func(context *gin.Context) {
		var service Service
		err := context.ShouldBindJSON(&service)
		if err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"error": "Failed to register service"})
			return
		}
		services = append(services, service)
		fmt.Printf("Registered service %s - %s: %d \n", service.Name, service.Host, service.Port)
		context.JSON(http.StatusOK, gin.H{"message": "Service registered successfully"})
	})
	router.GET("/discover/:name", func(context *gin.Context) {
		name := context.Param("name")
		for _, service := range services {
			if service.Name == name {
				context.JSON(http.StatusOK, gin.H{"host": service.Host, "port": service.Port})
				return
			}
		}
		context.JSON(http.StatusNotFound, gin.H{"error": "Service not found"})
	})
	err := router.Run(":8000")
	if err != nil {
		log.Fatal("Failed to start service-registry:", err)
	}
}
