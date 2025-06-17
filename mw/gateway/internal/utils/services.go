package utils

import (
	"context"
	"log"
	"strconv"

	srv "github.com/prabhakk-mw/learngo/mw/services/capitalize"
)

//grpcServerAddress, _ := microservices.StartMicroService(ctx)

/*
Microservices are of 2 kinds.
1. Single shot: Call and wait for it to complete.
2. Multi shot: Call, leave it running to serve future requests.
*/

func GetOrStartGRPCServer(ctx context.Context) {
	// TODO
}

func StartGRPCServer(ctx context.Context, server string) (serverAddress string, err error) {
	log.Printf("Starting GRPC Server for: %s", server)

	// In the future find a way to switch on the server string.
	// For now, just assume a service

	readyChan := make(chan int)
	errChan := make(chan error, 1)
	go srv.StartCapService(ctx, readyChan, errChan)

	// This line will block until the service responds with the port number.
	// It also marks that the service is ready to accept requests.
	select {
	case port := <-readyChan:
		portToUse := strconv.Itoa(port)
		grpcServerAddress := "localhost:" + portToUse
		return grpcServerAddress, nil

	case err := <-errChan:
		log.Println("Server failed to start:", err)
		return "", err
	case <-ctx.Done():
		err = ctx.Err()
		log.Println("Context cancelled while starting the capitalization service:", err)
		return "", err
	}

}
