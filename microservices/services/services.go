package services

import (
	"github.com/prabhakk-mw/microservices/service"
	"github.com/prabhakk-mw/microservices/services/environment"
)

func GetAvailableServices() []service.Service {

	return []service.Service{
		environment.GetServiceInfo(),
		// licensing.GetServiceInfo(),
		// session.GetServiceInfo(),
	}

}

// Think about a way where services can register themselves automatically
// This can be done using init functions in each service package
