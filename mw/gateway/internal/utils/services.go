package utils

import (
	"context"
	"log"

	"github.com/prabhakk-mw/learngo/mw/common/defs"
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
	serverInfoChan := make(chan defs.ServerInfo)
	errChan := make(chan error, 1)
	go srv.StartCapitalizeService(ctx, serverInfoChan, errChan)

	// This line will block until the service responds with server information.
	// It also marks that the service is ready to accept requests.
	select {
	case serverInfo := <-serverInfoChan:
		return serverInfo.GetAddress(), nil

	case err := <-errChan:
		log.Println("Server failed to start:", err)
		return "", err

	case <-ctx.Done():
		err := ctx.Err()
		log.Println("Context cancelled while starting the capitalization service:", err)
		return "", err
	}
}
