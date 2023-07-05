package main

import (
	"fmt"
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

func getProductServiceAddress() (string, error) {
	config := api.DefaultConfig()
	consulClient, err := api.NewClient(config)
	if err != nil {
		return "", err
	}

	entries, _, err := consulClient.Health().Service("product-service", "", true, nil)
	if err != nil {
		return "", err
	}

	if len(entries) == 0 {
		return "", fmt.Errorf("No healthy instances found for product-service")
	}

	address := fmt.Sprintf("%s:%d", entries[0].Service.Address, entries[0].Service.Port)
	return address, nil
}

func getOrderServiceAddress() (string, error) {
	config := api.DefaultConfig()
	consulClient, err := api.NewClient(config)
	if err != nil {
		return "", err
	}

	entries, _, err := consulClient.Health().Service("order-service", "", true, nil)
	if err != nil {
		return "", err
	}

	if len(entries) == 0 {
		return "", fmt.Errorf("No healthy instances found for order-service")
	}

	address := fmt.Sprintf("%s:%d", entries[0].Service.Address, entries[0].Service.Port)
	return address, nil
}

func main() {
	err := registerService("discovery-service", 8000)
	if err != nil {
		log.Fatal("Failed to register service:", err)
	}

	router := gin.Default()

	router.GET("/discover/:service", func(c *gin.Context) {
		service := c.Param("service")
		var address string
		var err error

		switch service {
		case "product-service":
			address, err = getProductServiceAddress()
		case "order-service":
			address, err = getOrderServiceAddress()
		default:
			c.JSON(http.StatusNotFound, gin.H{"error": "Service not found"})
			return
		}

		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"service": service, "address": address})
	})

	err = router.Run(":8000")
	if err != nil {
		log.Fatal("Failed to start discovery-service:", err)
	}
}
