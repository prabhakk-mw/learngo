package environment

import "github.com/prabhakk-mw/microservices/service"

var version = service.Version{Major: 1, Minor: 0, Patch: 0}

var environmentService = service.Service{
	Name:        "Environment",
	Description: "Provides information about the environment",
	Endpoint:    "/environment",
	Port:        8081,
	HealthCheck: "/health",
	Version:     version,
}

func GetServiceInfo() service.Service {
	return environmentService
}

func GetEnvironmentInfo() (environment string) {
	environment = "This is a Windows Environment"
	return
}
