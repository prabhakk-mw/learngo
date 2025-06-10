package environment

import (
	"fmt"

	"github.com/prabhakk-mw/microservices/service"
)

var version = service.Version{Major: 1, Minor: 0, Patch: 0}

var port = 8081

func environmentStarter() service.ServiceResponse {
	fmt.Println("Starting Environment Service...")
	return service.ServiceResponse{
		Port:   port,
		Status: "Started",
	}
}

var environmentService = service.Service{
	Name:          "Environment",
	Description:   "Provides information about the environment",
	HTTPEndpoints: []string{"/start", "/stop", "/status", "/info"},
	Version:       version,
	Start:         environmentStarter,
}

func GetServiceInfo() service.Service {
	return environmentService
}
