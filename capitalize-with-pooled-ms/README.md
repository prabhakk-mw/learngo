# Capitalize + Spawn Microservice in different MODULES + using a Pooled Microservice
[WIP]

This evolution of the repository is based on the base implementation in "../capitalize-ms-in-another-module"

The primary difference here is that we will attempt to start the microservice from within the http server 
**and** *the microservice will be provided by a different MODULE*
**and** *a new microservice will not be started for each request. Instead, if one has not been started, a new one will be started, and kept alive until program shutsdown.*

This evolution showcases that it is possible to keep microservices alive to serve multiple requests.


## How to use

```bash
# Navigate to folder with go.mod
$ cd capitalize
# Get dependencies to run protoc
$ go install google.golang.org/protobuf/cmd/protoc-gen-go@latest && go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
# Run protoc
$ protoc --go_out=./proto --go-grpc_out=./proto proto/capservice.proto
# Open two terminals and run the two servers
## Terminal 1
$ go run http-server/main.go
## Terminal 2
$ go run capservice/main.go
```

Then open a browser, and navigate to `localhost:8081/grpc?payload=stringToCapitalize`

And update query parameter value to view the grpc service execution.

Other supported URLs are

`localhost:8081/stringToCapitalize`
`localhost:8081/caps&payload=stringToCapitalize`



## Phase 1

* HTTP Server & Cap Service are started manually and independently


## Evolutions

* The HTTP Server accepts gRPC traffic?
* The Cap Service is its own Go Module
* The Cap Service is a executable that is started by the HTTP Server.



## Things to try:

When an HTTP listener in Go shuts down, a message can be printed to indicate the shutdown process. This is typically achieved by handling signals, such as SIGINT (Ctrl+C) or SIGTERM, and then gracefully shutting down the server. 
Here's how it generally works: 

• Signal Handling: The os/signal package is used to listen for system signals that indicate a shutdown request. 
• Context: A context.Context is often used to manage the shutdown process. This context can be canceled when a shutdown signal is received. 
• Server Shutdown: The http.Server's Shutdown method is called with the context. This method closes all open listeners, closes idle connections, and waits for all connections to become idle before shutting down. 
• Logging: A message is printed before and after the Shutdown call to indicate the start and completion of the shutdown process. 

Here's a basic example: 
```go
package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	// Create a new server
	srv := &http.Server{Addr: ":8080", Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello, World!")
	})}

	// Create a channel to receive signals
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)

	// Start the server in a goroutine
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("HTTP server error: %v", err)
		}
	}()

	// Wait for a shutdown signal
	sig := <-sigChan
	log.Printf("%v signal received, shutting down...", sig)

	// Create a context with a timeout for the shutdown
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Shutdown the server
	if err := srv.Shutdown(ctx); err != nil {
		log.Printf("HTTP server shutdown error: %v", err)
	}

	log.Println("HTTP server shutdown complete.")
}
```
In this example: 

• When the server starts, it listens on port 8080. 
• The signal.Notify function sets up the program to receive SIGINT (Ctrl+C) and SIGTERM signals. 
• When the program receives one of these signals, it initiates the shutdown process. 
• The server is gracefully shut down, and a message is printed to indicate completion. 
• If there is a problem during shutdown, an error message will be printed. 

AI responses may include mistakes.

[-] https://mohanliu.com/blog/Design%20Discuss%20and%20Re-invent%20a%20go%20routine%20Container[-] https://github.com/bnixon67/go-weblogin
