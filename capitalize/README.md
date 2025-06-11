# Capitalize

* A HTTP server that accepts text as input using REST endpoints
* The server then uses a capitalization service to convert the text to Upper Case.
* Returns the capitalized Upper case response.

The communication between the HTTP Server and the Capitalization service is through gRPC.

Both the HTTP Server & Cap Service as a part of a single GO Module.
The Cap Service is implemented as a package.

## Phase 1

* HTTP Server & Cap Service are started manually and independently


## Evolutions

* The HTTP Server accepts gRPC traffic?
* The Cap Service is its own Go Module
* The Cap Service is a executable that is started by the HTTP Server.


## Thoughts

* What is the point of gRPC communication, when you can simply do a function call to make the communication possible?


Preference of communication patterns:
1. Simple function call
2. 


