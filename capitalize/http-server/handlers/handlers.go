package handlers

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strings"
	"time"

	pb "github.com/prabhakk-mw/learngo/capitalize/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	grpcPort          = ":50051"
	grpcServerAddress = "localhost" + grpcPort
)

func BasicHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Received request:", r.Method, r.URL.Path)
	payload := string(r.URL.Path[1:])
	fmt.Fprintf(w, "Hi there, I love %s!", strings.ToUpper(payload))
}

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
		conn, err := grpc.NewClient(grpcServerAddress, grpc.WithTransportCredentials(insecure.NewCredentials()))
		if err != nil {
			http.Error(w, "Failed to connect to gRPC server", http.StatusInternalServerError)
			return
		}
		defer conn.Close()

		client := pb.NewCapServiceClient(conn)
		capitalizedPayload, err := callCapGrpcService(client, payload)
		if err != nil {
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
