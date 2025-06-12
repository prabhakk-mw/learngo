package microservices

import (
	"context"
	"log"
	"strconv"

	"github.com/prabhakk-mw/learngo/capitalize-spawn-ms/capservice"
	"github.com/prabhakk-mw/learngo/capitalize-spawn-ms/http-server/ports"
)

func findNextFreePort() (int, error) {
	// This function is a placeholder. In a real application, you would implement logic to find the next free port.
	port, err := ports.GetFreePort()
	if err != nil {
		log.Println("Failed to find a free port:", err)
		return 0, err
	}
	return port, nil
}

func StartMicroService(ctx context.Context) (grpcServerAddress string, err error) {
	port, err := findNextFreePort()
	if err != nil {
		return "", err
	}

	portToUse := strconv.Itoa(port)
	grpcServerAddress = "localhost:" + portToUse
	// errChan := make(chan error)
	// go capservice.StartCapServiceOn(":"+portToUse, errChan)
	go capservice.StartCapService(ctx, ":"+portToUse)
	// if err := <-errChan; err != nil {
	// 	log.Println("Failed to start the capitalization service:", err)
	// } else {
	// 	log.Printf("Capitalize Microservice started on %s\n", grpcServerAddress)
	// }
	log.Printf("Capitalize Microservice started on %s\n", grpcServerAddress)
	return grpcServerAddress, err
}
