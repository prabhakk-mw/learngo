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

	readyChan := make(chan int)
	doneChan := make(chan int)
	go capservice.TestService(ctx, readyChan, doneChan)

	// wait here until the servie is ready.
	log.Println("Execution blocked until ready signal is received.")
	readySignal := <-readyChan
	log.Printf("TestService is ready to accept requests, : %d\n", readySignal)

	// Wait here until service / context is done
	select {
	case <-doneChan:
		log.Println("TestService completed successfully")
	case <-ctx.Done():
		log.Println("Main context done, exiting...")
	}
}
