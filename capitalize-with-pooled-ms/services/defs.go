package services

import "context"

type Service struct {
	Name        string
	Port        int
	Description string
}

type Controller interface {
	// Start starts the service.
	StartService(ctx context.Context, port int, ready chan<- struct{}, err chan<- error)
}
