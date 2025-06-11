package main

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

func callCapGrpcService(client pb.CapServiceClient, payload string) (capitalizedPayload string, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	res, err := client.Capitalize(ctx, &pb.CapRequest{Payload: payload})
	if err != nil {
		log.Fatalf("could not call Capitalize: %v", err)
		return "", err
	}

	fmt.Printf("Response from gRPC service: %s\n", res.GetPayload())
	capitalizedPayload = res.GetPayload()
	return capitalizedPayload, nil
}

func grpcHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("gRPC request received")

	u, err := url.Parse(r.URL.String())
	if err != nil {
		http.Error(w, "Error parsing URL", http.StatusBadRequest)
		return
	}

	// Extract query parameters
	query := u.Query()

	// Get single value
	payload := query.Get("payload")

	conn, err := grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
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
	fmt.Printf("Capitalized payload: %s\n", capitalizedPayload)
	// Write the response to the HTTP client
	fmt.Fprintf(w, "Capitalized %s to %s \n", payload, capitalizedPayload)
}

func localHandler(w http.ResponseWriter, r *http.Request) {
	// Parse the URL
	u, err := url.Parse(r.URL.String())
	if err != nil {
		http.Error(w, "Error parsing URL", http.StatusBadRequest)
		return
	}

	// Extract query parameters
	query := u.Query()

	// Get single value
	param1 := query.Get("string")
	fmt.Fprintf(w, "string: %s\n", strings.ToUpper(param1))
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Received request:", r.Method, r.URL.Path)
	payload := string(r.URL.Path[2:])
	fmt.Fprintf(w, "Hi there, I love %s!", strings.ToUpper(payload))
}

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/caps/", localHandler)
	http.HandleFunc("/grpc/", grpcHandler)
	log.Fatal(http.ListenAndServe(":8081", nil))
}
