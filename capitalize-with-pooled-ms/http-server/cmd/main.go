package main

import (
	"context"
	"log"

	"github.com/prabhakk-mw/learngo/capitalize-ms-in-another-module/capservice"
)

const (
	httpPort = ":8081"
)

// func old_main() {
// 	log.Printf("Started Web Server on localhost%s\n", httpPort)
// 	log.Printf("Use : http://localhost%s/grpc?payload=yourtext to capitalize text\n", httpPort)
// 	http.HandleFunc("/", handlers.QueryParamHandler)
// 	http.HandleFunc("/grpc/", handlers.GRPCHandler)
// 	log.Fatal(http.ListenAndServe(httpPort, nil))
// }

func main() {

	// Calling test service as a goroutine
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	readyChan := make(chan struct{})
	doneChan := make(chan struct{})
	go capservice.TestService(ctx, readyChan, doneChan)

	select {
	case <-readyChan:
		log.Println("TestService is ready to accept requests")
	case <-doneChan:
		log.Println("TestService completed successfully")
	case <-ctx.Done():
		log.Println("Main context done, exiting...")
	}
}
