package main

import (
	"log"
	"net/http"

	"github.com/prabhakk-mw/learngo/mw/gateway/internal/handlers"
)

const (
	httpPort = ":8081"
)

func main() {
	log.Printf("Started Web Server on localhost%s\n", httpPort)

	http.HandleFunc("/capitalize/", handlers.CapitalizeHandler)
	log.Printf("Use : http://localhost%s/capitalize?payload=yourtext to capitalize text\n", httpPort)

	log.Fatal(http.ListenAndServe(httpPort, nil))
}
