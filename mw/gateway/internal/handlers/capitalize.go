package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"net/url"

	"github.com/prabhakk-mw/learngo/mw/common/defs"

	pb "github.com/prabhakk-mw/learngo/mw/services/capitalize/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// type defs.Handlers struct {
// 	RootCtx        context.Context
// 	grpcServerInfo *defs.ServerInfo
// }

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
func (handlers *defs.Handlers) capitalizeHandler(w http.ResponseWriter, r *http.Request, reuseServer bool) {

	if payload := getPayload(r); len(payload) != 0 {

		// Function call style.
		// capitalize.CapitalizeCore(payload)

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

		res, err := callCapGrpcService(client, payload)
		if err != nil {
			log.Printf("Failed to call gRPC service: %v\n", err)
			http.Error(w, "Failed to call gRPC service", http.StatusInternalServerError)
			return
		}
		// Write the response to the HTTP client
		w.Header().Set("Content-Type", "application/json")

		// Encode the response object to JSON and write to the response writer
		if err := json.NewEncoder(w).Encode(res); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
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
func (handlers *defs.Handlers) CapitalizeHandler(w http.ResponseWriter, r *http.Request) {
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
func (handlers *defs.Handlers) StaticCapitalizeHandler(w http.ResponseWriter, r *http.Request) {
	handlers.capitalizeHandler(w, r, true)
}
