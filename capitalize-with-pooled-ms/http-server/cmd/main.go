package main

import (
	"log"
	"net/http"

	"github.com/prabhakk-mw/learngo/capitalize-ms-in-another-module/http-server/handlers"
)

const (
	httpPort = ":8081"
)

func main() {
	log.Printf("Started Web Server on localhost%s\n", httpPort)
	log.Printf("Use : http://localhost%s/grpc?payload=yourtext to capitalize text\n", httpPort)
	http.HandleFunc("/", handlers.QueryParamHandler)
	http.HandleFunc("/grpc/", handlers.GRPCHandler)
	log.Fatal(http.ListenAndServe(httpPort, nil))
}
