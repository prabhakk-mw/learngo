package licensing

import "github.com/prabhakk-mw/microservices/service"

var version = service.Version{Major: 1, Minor: 0, Patch: 0}

var licensingService = service.Service{
	Name:        "Licensing",
	Description: "Provides information about the licensing",
	Endpoints:   []string{"/start", "/stop", "/status", "/info"},
	Port:        8082,
	HealthCheck: "/health",
	Version:     version,
}

func GetServiceInfo() service.Service {
	return licensingService
}

func GetLicensingInfo() (licensing string) {
	licensing = "This is a Windows Licensing"
	return
}
