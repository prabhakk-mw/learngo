package handlers

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"time"

	"github.com/prabhakk-mw/learngo/mw/common/defs"
	"github.com/prabhakk-mw/learngo/mw/gateway/internal/utils"
	pb "github.com/prabhakk-mw/learngo/mw/services/capitalize/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Handlers struct {
	RootCtx        context.Context
	grpcServerInfo *defs.ServerInfo
}

/***** Local Functions *****/

func getPayload(r *http.Request) string {

	u, err := url.Parse(r.URL.String())
	if err != nil {
		return ""
	}
	// Extract query parameters
	query := u.Query()
	// Get single value
	return query.Get("payload")
}

func callCapGrpcService(client pb.CapServiceClient, payload string) (capitalizedPayload string, err error) {

	// Create a context that expects the grpc server to respond within 1 second.
	grpcCallCtx, grpcCallCancel := context.WithTimeout(context.Background(), time.Second)
	defer grpcCallCancel()

	res, err := client.Capitalize(grpcCallCtx, &pb.CapRequest{Payload: payload})
	if err != nil {
		log.Printf("Could not call Capitalize: %v\n", err)
		return "", err
	}

	capitalizedPayload = res.GetPayload()
	return capitalizedPayload, nil
}

func (handlers *Handlers) startGRPCServer(reuseServer bool) (serverInfo defs.ServerInfo, cancel context.CancelFunc) {

	if reuseServer {
		if handlers.grpcServerInfo == nil {
			// The lifetime of the reusable server is tied to the root context.
			serverInfo = utils.StartGRPCServer(handlers.RootCtx, "capitalize")
			handlers.grpcServerInfo = &serverInfo
		} else {
			serverInfo = *handlers.grpcServerInfo
		}
		cancel = nil

	} else {
		// The context is per request now, this ensures that the server is shutdown when the request is complete.
		ctx, cancelFcn := context.WithTimeout(handlers.RootCtx, 3*time.Second)
		serverInfo = utils.StartGRPCServer(ctx, "capitalize")
		cancel = cancelFcn
	}
	return serverInfo, cancel

}

/*
This handler has a need to call a gRPC service to capitalize the payload.
To do so, it needs to know:
1. which microservice can serve its needs
2. which endpoint on the microservice to call

Ideas:
1. If the microservice is not running, start it, and cache it for future requests.
2. Let the client of the http server be responsible for passing in the name of the microservice it needs?
  - This may be a nice way to dovetail into the gateway concept?

3. Think about a service discovery mechanism to find the available microservices?
  - May not be needed if the main server hardcodes the services that it needs!

4. Move the microservice to another module,
*/
func (handlers *Handlers) capitalizeHandler(w http.ResponseWriter, r *http.Request, reuseServer bool) {

	if payload := getPayload(r); len(payload) != 0 {

		grpcServerInfo, cancel := handlers.startGRPCServer(reuseServer)
		if cancel != nil {
			defer cancel()
		}

		conn, err := grpc.NewClient(grpcServerInfo.GetAddress(), grpc.WithTransportCredentials(insecure.NewCredentials()))
		if err != nil {
			http.Error(w, "Failed to connect to gRPC server", http.StatusInternalServerError)
			return
		}
		defer conn.Close()

		client := pb.NewCapServiceClient(conn)

		capitalizedPayload, err := callCapGrpcService(client, payload)
		if err != nil {
			log.Printf("Failed to call gRPC service: %v\n", err)
			http.Error(w, "Failed to call gRPC service", http.StatusInternalServerError)
			return
		}
		// Write the response to the HTTP client
		fmt.Fprintf(w, "Capitalized [%s] to [%s] \n", payload, capitalizedPayload)
	} else {
		// Even though the service allows for null strings, this simulates business logic that does not accept null strings.
		http.Error(w, "Error parsing URL.", http.StatusBadRequest)
	}
}

/***** Local Functions *****/

// CapitalizeHandler Capitalize strings using
//
//	@Summary		Uses a grpc server per request
//	@Description	Capitalize URL Query Parameters using gRPC service as a microservice
//	@Tags			Capitalize Example
//	@Accept			json
//	@Produce		json
//	@Param			payload	query		string		true	"String to Capitalize"
//	@Success		200		{string}	Helloworld	"Capitalized String"
//	@Router			/capitalize [get]
func (handlers *Handlers) CapitalizeHandler(w http.ResponseWriter, r *http.Request) {
	handlers.capitalizeHandler(w, r, false)
}

// CapitalizeHandler Capitalize strings using Static gRPC Handler
//
//	@Summary		Uses the same grpc server for the lifetime of the program
//	@Description	Capitalize URL Query Parameters using gRPC service as a microservice
//	@Tags			Capitalize Example
//	@Accept			json
//	@Produce		json
//	@Param			payload	query		string		true	"String to Capitalize"
//	@Success		200		{string}	Helloworld	"Capitalized String"
//	@Router			/static-capitalize [get]
func (handlers *Handlers) StaticCapitalizeHandler(w http.ResponseWriter, r *http.Request) {
	handlers.capitalizeHandler(w, r, true)
}
