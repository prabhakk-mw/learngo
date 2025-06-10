# Microservices using gRPC written in GO

## Goals
1. Create a microservice based project
2. Demonstrate that a single executable can be used to spin up multiple microservices to perform work on demand
3. All communication between the microservices must be via gRPC
4. All communication between microservices must be socket based. 

## Introduction

In this repository, we will create an executable called `driver`

The `driver` executable, when run will start operating on a PORT

This will allow external programs (a.k.a `clients`) to communicate with it using TCP using REST endpoints.
Additionally, `clients` can also communicate with the `driver` via a gRPC interface.

The `driver` program will provide access to some capabilities/functionality.

These capabilities will be served by other microservices.

In this example, `driver` will enable access to the following services to use a software called `MATLAB`

1. Environment Service: Query PATH and FileSystem to provide location at which MATLAB is installed.
2. Licensing Service: Accept login credentials, to authenticate/authorize use of MATLAB
3. Session Service: Start a MATLAB session and provide connection details.
    1. Execution Interface: Communicate with this MATLAB to perform tasks.


### Stage1: Setup a driver program