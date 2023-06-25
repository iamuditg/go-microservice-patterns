package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func main() {
	router := gin.Default()
	router.GET("/products", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "List of products"})
	})

	err := router.Run(":8002")
	if err != nil {
		log.Fatal("Failed to start product-service:", err)
	}
}
