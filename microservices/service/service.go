package service

import "fmt"

type ServiceInfo interface {
	GetServiceInfo() Service
}

type Version struct {
	Major int
	Minor int
	Patch int
}

func (v Version) String() string {
	return fmt.Sprintf("%d.%d.%d", v.Major, v.Minor, v.Patch)
}

type ServiceResponse struct {
	Port   int
	Status string
}

type ServiceController interface {
	Start() ServiceResponse
}

type ServiceStarter func() ServiceResponse

type Service struct {
	Name          string
	Description   string
	HTTPEndpoints []string
	Port          int
	Version       Version
	Start         ServiceStarter
}

func (service Service) String() string {
	return fmt.Sprintf("Service Name: %s\n Description: %s\n Endpoint: %s\n Port: %d\n Version: %s\n", service.Name, service.Description, service.HTTPEndpoints, service.Port, service.Version.String())
}
