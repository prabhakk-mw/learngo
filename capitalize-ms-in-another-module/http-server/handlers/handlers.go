package handlers

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strings"
	"time"

	pb "github.com/prabhakk-mw/learngo/capitalize-ms-in-another-module/capservice/pb"
	"github.com/prabhakk-mw/learngo/capitalize-ms-in-another-module/http-server/microservices"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func QueryParamHandler(w http.ResponseWriter, r *http.Request) {
	// Parse the URL
	u, err := url.Parse(r.URL.String())
	if err != nil {
		http.Error(w, "Error parsing URL", http.StatusBadRequest)
		return
	}

	// Extract query parameters
	query := u.Query()
	if len(query) != 0 {
		// Get single value
		payload := query.Get("payload")
		capitalizedPayload := strings.ToUpper(payload)
		fmt.Fprintf(w, "Capitalized %s to %s \n", payload, capitalizedPayload)
	} else {
		fmt.Fprintf(w, "No query parameters provided. Please use ?payload=your_text to capitalize.\n")
	}
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
func GRPCHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("gRPC request received")

	u, err := url.Parse(r.URL.String())
	if err != nil {
		http.Error(w, "Error parsing URL", http.StatusBadRequest)
		return
	}

	// Extract query parameters
	query := u.Query()

	// Get single value
	payload := query.Get("payload")
	if len(payload) != 0 {

		log.Println("About to start microservice")
		// Using the context.WithCancel sometimes gives the "context deadline exceeded" error,
		// Using WithTimeout to give atleast 10 seconds for the microservice to start
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		// Without this, the microservice would have to be started manually
		grpcServerAddress, _ := microservices.StartMicroService(ctx)
		log.Println("Started microservice at address:", grpcServerAddress)

		conn, err := grpc.NewClient(grpcServerAddress, grpc.WithTransportCredentials(insecure.NewCredentials()))
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
		// Write the response
		log.Printf("Capitalized payload: %s\n", capitalizedPayload)
		// Write the response to the HTTP client
		fmt.Fprintf(w, "Used gRPC service to capitalize %s to %s \n", payload, capitalizedPayload)
	} else {
		fmt.Fprintf(w, "No query parameters provided. Please use ?payload=your_text to capitalize.\n")
	}
}

func callCapGrpcService(client pb.CapServiceClient, payload string) (capitalizedPayload string, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	res, err := client.Capitalize(ctx, &pb.CapRequest{Payload: payload})
	if err != nil {
		log.Printf("Could not call Capitalize: %v\n", err)
		return "", err
	}

	capitalizedPayload = res.GetPayload()
	log.Printf("Response from gRPC service: %s\n", capitalizedPayload)
	return capitalizedPayload, nil
}
