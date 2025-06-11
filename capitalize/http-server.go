package main

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strings"
)

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
	log.Fatal(http.ListenAndServe(":8081", nil))
}
