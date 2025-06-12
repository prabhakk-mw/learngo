package main

import (
	"log"
	"net/http"

	"github.com/prabhakk-mw/learngo/capitalize/http-server/handlers"
)

const (
	httpPort = ":8081"
)

func main() {
	log.Printf("Started Web Server on localhost%s\n", httpPort)
	http.HandleFunc("/", handlers.BasicHandler)
	http.HandleFunc("/caps/", handlers.QueryParamHandler)
	http.HandleFunc("/grpc/", handlers.GRPCHandler)
	log.Fatal(http.ListenAndServe(httpPort, nil))
}
