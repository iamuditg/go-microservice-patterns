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
	err := registerService("order-service", 8002)
	if err != nil {
		log.Fatal("Failed to register service:", err)
	}

	router := gin.Default()

	router.POST("/orders", func(c *gin.Context) {
		// Logic to add an order
		// ...

		// Example response
		order := map[string]interface{}{
			"orderId":    1,
			"product":    "Product 1",
			"quantity":   2,
			"totalPrice": 100,
		}
		c.JSON(http.StatusOK, order)
	})

	err = router.Run(":8002")
	if err != nil {
		log.Fatal("Failed to start order-service:", err)
	}
}
