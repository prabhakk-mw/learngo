package main

import (
	"fmt"

	"github.com/prabhakk-mw/microservices/service"
	"github.com/prabhakk-mw/microservices/services"
)

func main() {

	availableServices := services.GetAvailableServices()

	var environmentService *service.Service
	for _, service := range availableServices {
		if service.Name == "Environment" {
			environmentService = &service
			break
		}
	}

	fmt.Println(environmentService)

	envServiceInfo := environmentService.Start()
	fmt.Printf("Service %s started on port %d with status: %s\n", environmentService.Name, envServiceInfo.Port, envServiceInfo.Status)

	// At this point the service has started on the returned port.
	// Then this program can start communicating with this service via gRPC

	// So aprt rom adhering to the interfaces specified by service.go
	// The services will also have communicate their gRPC protobufs!
	// How do we establish that..

	// Finally, how will all this look when its between modules.
	// Its all fun and games when the service is a package in the same module.
	// Now if its between modules,

	// Then the service module, will have to fetch the service.go package.
	// instantiate itself,
	// register with the service registry,
	//

}
