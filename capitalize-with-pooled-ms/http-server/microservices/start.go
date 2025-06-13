package microservices

import (
	"context"
	"log"
	"strconv"

	"github.com/prabhakk-mw/learngo/capitalize-ms-in-another-module/capservice"
	"github.com/prabhakk-mw/learngo/capitalize-ms-in-another-module/http-server/ports"
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

	readyChan := make(chan struct{})
	errChan := make(chan error, 1)
	go capservice.StartCapService(ctx, portToUse, readyChan, errChan)

	select {
	case <-readyChan:
		log.Println("Server is ready")
		return grpcServerAddress, nil

		// Continue startup
	case err := <-errChan:
		log.Println("Server failed to start:", err)
		return "", err
	case <-ctx.Done():
		err = ctx.Err()
		log.Println("Context cancelled while starting the capitalization service:", err)
		return "", err
	}
}
