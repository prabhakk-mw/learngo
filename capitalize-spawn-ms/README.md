# Capitalize + Spawn Microservice

This evolution of the repository is based on the base implementation in "../capitalize"

The primary difference here is that we will attempt to start the microservice from within the http server.
Previously, one was required to start the microservice independantly using `go run capitalize/main.go`


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



