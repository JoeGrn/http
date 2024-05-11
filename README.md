# HTTP

A http server built in Go that's capable of handling GET/POST requests, serving files, gzip compression and handling multiple concurrent connections.

## Usage

`go build` - build the application this will output and executable 'http'

`./http` - run the http server

`curl -v GET http://localhost:4221/` - get a 200 response

`curl -v GET http://localhost:4221/notfound` - get a 404 response

`curl -v http://localhost:4221/echo/hello` - get a response body

`curl -v http://localhost:4221/user-agent` - get the request headers

`curl -v GET http://localhost:4221/files/non-existent` - get a file (need to create the folder first)

`curl -v POST http://localhost:4221/files/README.md` - post a file