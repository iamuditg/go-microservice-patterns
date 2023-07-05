# Service Discovery using Consul in Golang

This repository contains an example implementation of service discovery using Consul in Golang.

## Introduction

Service discovery is a critical aspect of distributed systems where services need to dynamically discover and communicate with each other. Consul, a popular service mesh and service discovery tool, provides a simple and reliable way to register, discover, and monitor services in a distributed environment.

This example consists of three services: product-service, order-service, and discovery-service. The product-service and order-service are the actual services that provide specific functionalities, while the discovery-service acts as a middleware for service discovery.

## Features

The implemented example includes the following features:

- **Product Service**: Exposes an API endpoint to retrieve a list of products.

- **Order Service**: Exposes an API endpoint to add an order.

- **Discovery Service**: Provides an API endpoint to discover the address of a specified service.

Additionally, the implementation includes:

- **Consul Integration**: Utilizes the Consul client library for service registration, health checks, and more advanced features.

- **Error Handling**: Includes error handling to handle service discovery failures and invalid requests.

- **Configuration Management**: Configures the services to register with Consul and communicate with each other.

- **Logging**: Logs important events and errors for better visibility and troubleshooting.

## Prerequisites

To run the example, make sure you have Consul installed and running in your environment. Follow the official Consul documentation for installation instructions.

## Usage

1. Start the product-service:
   ```bash
   cd product-service
   go run main.go

2. Start the order-service:
   ```bash
   cd order-service
   go run main.go

3. Start the discovery-service:
   ```bash
   cd discovery-service
   go run main.go


Conclusion
Service discovery is a crucial component of microservices architecture, enabling dynamic and resilient communication between services. Consul provides a powerful solution for service discovery and offers additional features like health checks, key-value store, and distributed locking.

This example demonstrates how to implement service discovery using Consul in Golang, enabling seamless communication between services and enhancing the reliability and scalability of your distributed systems.
