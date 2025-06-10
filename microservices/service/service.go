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

type Service struct {
	Name        string
	Description string
	Endpoint    string
	Port        int
	HealthCheck string
	Version     Version
}
