package session

import "github.com/prabhakk-mw/microservices/service"

var version = service.Version{Major: 1, Minor: 0, Patch: 0}

var sessionService = service.Service{
	Name:        "Session",
	Description: "Provides information about the session",
	Endpoints:   []string{"/start", "/stop", "/status", "/info"},
	Port:        8083,
	HealthCheck: "/health",
	Version:     version,
}

func GetServiceInfo() service.Service {
	return sessionService
}

func GetSessionInfo() (session string) {
	session = "This is a Windows Session"
	return
}
