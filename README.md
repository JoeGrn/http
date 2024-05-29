# HTTP

A http server built in Go that's capable of handling GET/POST requests, serving files, gzip compression and handling multiple concurrent connections.

## Requirements

* `go` v1.22

## Usage

`go build ./cmd/http` - build the application this will output an executable 'http'

`./http` - run the http server

`docker build -t http .` - build the application using docker if you do not have go on your system

`docker run -p 8080:8080 http` - run the container exposing port 8080

`curl -v GET http://localhost:4221/` - get a 200 response

`curl -v GET http://localhost:4221/notfound` - get a 404 response

`curl -v http://localhost:4221/echo/hello` - get a response body

`curl -v http://localhost:4221/user-agent` - get the request headers

`curl -v GET http://localhost:4221/files/non-existent` - get a file (need to create the folder first)

`curl -v POST http://localhost:4221/files/README.md` - post a file