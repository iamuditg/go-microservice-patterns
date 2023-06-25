package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func main() {
	router := gin.Default()
	router.GET("/orders", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{"message": "List of orders"})
	})
	err := router.Run(":8001")
	if err != nil {
		log.Fatal("Failed to start order-services:", err)
	}
}
