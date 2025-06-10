package main

import (
	"fmt"

	"github.com/prabhakk-mw/microservices/services"
	"github.com/prabhakk-mw/microservices/services/environment"
)

func main() {

	fmt.Println(environment.GetEnvironmentInfo())
	availableServices := services.GetAvailableServices()

	for _, service := range availableServices {
		fmt.Printf("Service Name: %s\n", service.Name)
		fmt.Printf("Description: %s\n", service.Description)
		fmt.Printf("Endpoint: %s\n", service.Endpoint)
		fmt.Printf("Port: %d\n", service.Port)
		fmt.Printf("Health Check: %s\n", service.HealthCheck)
		fmt.Printf("Version: %s\n", service.Version.String())
		fmt.Println()
	}
}
