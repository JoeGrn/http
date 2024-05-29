# HTTP

A http server built in Go that's capable of handling GET/POST requests, serving files, gzip compression and handling multiple concurrent connections.

## Requirements

`go` v1.22

## Makefile

To simplify the build and run process, you can use the following Makefile commands:

`make all` - test and build the application

`make build` - build the application and output the executable into /dist

`make run` - build and run the application

`make fmt` - format the codebase with go fmt

`make test` - run the unit tests

## Docker

To build and run the HTTP server using Docker, you can use the following commands:

`make docker-build` - build the docker image

`make docker-run` - run the image exposed on port 8080

## Usage

`curl -v GET http://localhost:4221/` - get a 200 response

`curl -v GET http://localhost:4221/notfound` - get a 404 response

`curl -v http://localhost:4221/echo/hello` - get a response body

`curl -v http://localhost:4221/user-agent` - get the request headers

`curl -v GET http://localhost:4221/files/non-existent` - get a file (need to create the folder first)

`curl -v POST http://localhost:4221/files/README.md` - post a file