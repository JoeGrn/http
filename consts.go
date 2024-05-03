package main

const (
	PROTOCOL              = "HTTP/1.1 "
	HTTP_STATUS_OK        = "200 OK"
	HTTP_STATUS_CREATED   = "201 Created"
	HTTP_STATUS_NOT_FOUND = "404 Not Found"
	HTTP_STATUS_BAD_REQ   = "400 Bad Request"
	HTTP_STATUS_INTERNAL  = "500 Internal Server Error"
	HTTP_STATUS_NOT_IMPL  = "501 Not Implemented"
	SEPARATOR             = "\r\n"
)

const (
	CONTENT_TYPE_HEADER   = "Content-Type: "
	CONTENT_LENGTH_HEADER = "Content-Length: "
)

const (
	CONTENT_TYPE_TEXT   = "text/plain"
	CONTENT_TYPE_STREAM = "application/octet-stream"
)

const (
	FILES_PATH      = "/files/"
	ECHO_PATH       = "/echo/"
	USER_AGENT_PATH = "/user-agent"
)

const (
	GET    = "GET"
	POST   = "POST"
	PUT    = "PUT"
	DELETE = "DELETE"
)
