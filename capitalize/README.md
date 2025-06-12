# Capitalize

* A HTTP server that accepts text as input using REST endpoints
* The server then uses a capitalization service to convert the text to Upper Case.
* Returns the capitalized Upper case response.

The communication between the HTTP Server and the Capitalization service is through gRPC.

Both the HTTP Server & Cap Service as a part of a single GO Module.
The Cap Service is implemented as a package.


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



