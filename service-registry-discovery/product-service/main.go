package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hashicorp/consul/api"
)

func registerService(serviceName string, port int) error {
	config := api.DefaultConfig()
	consulClient, err := api.NewClient(config)
	if err != nil {
		return err
	}

	registration := new(api.AgentServiceRegistration)
	registration.Name = serviceName
	registration.Port = port

	err = consulClient.Agent().ServiceRegister(registration)
	if err != nil {
		return err
	}

	return nil
}

func main() {
	err := registerService("product-service", 8001)
	if err != nil {
		log.Fatal("Failed to register service:", err)
	}

	router := gin.Default()

	router.GET("/products", func(c *gin.Context) {
		// Logic to fetch products
		// ...

		// Example response
		products := []string{"Product 1", "Product 2", "Product 3"}
		c.JSON(http.StatusOK, products)
	})

	err = router.Run(":8001")
	if err != nil {
		log.Fatal("Failed to start product-service:", err)
	}
}
