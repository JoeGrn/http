package main

import (
	"fmt"
	"net"
	"os"
	"path/filepath"
	"strings"
)

func main() {

	l, err := net.Listen("tcp", "0.0.0.0:4221")

	if err != nil {
		fmt.Println("Failed to bind to port 4221")
		os.Exit(1)
	}

	defer l.Close()

	for {
		conn, err := l.Accept()
		if err != nil {
			fmt.Println("Error accepting connection: ", err.Error())
			os.Exit(1)
		}

		go handleConnection(conn)
	}

}

func handleConnection(conn net.Conn) {
	defer conn.Close()
	buf := make([]byte, 1024)
	for {
		size, err := conn.Read(buf)
		if err != nil {
			fmt.Println("Error reading connection: ", err.Error())
			return
		}
		request := string(buf[:size])
		res := handleResponse(request)

		conn.Write([]byte(res))
	}
}

func handleResponse(request string) string {
	req := ParseRequest(request)

	switch {
	case req.path == "/":
		response := NewResponse()
		response.SetProtocol(PROTOCOL)
		response.SetStatus(HTTP_STATUS_OK)
		return response.String()
	case strings.HasPrefix(req.path, ECHO_PATH):
		body := strings.ReplaceAll(req.path, ECHO_PATH, "")

		response := NewResponse()
		response.SetProtocol(PROTOCOL)
		response.SetStatus(HTTP_STATUS_OK)

		if strings.Contains(req.headers[ACCEPT_ENCODING_HEADER], ENCODING_TYPE_GZIP) {
			response.SetHeader(CONTENT_ENCODING_HEADER, ENCODING_TYPE_GZIP)
		}

		response.SetHeader(CONTENT_TYPE_HEADER, CONTENT_TYPE_TEXT)
		response.SetHeader(CONTENT_LENGTH_HEADER, fmt.Sprintf("%d", len(body)))
		response.SetBody(body)

		return response.String()
	case strings.HasPrefix(req.path, USER_AGENT_PATH):
		response := NewResponse()
		response.SetProtocol(PROTOCOL)
		response.SetStatus(HTTP_STATUS_OK)
		response.SetHeader(CONTENT_TYPE_HEADER, CONTENT_TYPE_TEXT)
		response.SetHeader(CONTENT_LENGTH_HEADER, fmt.Sprintf("%d", len(req.headers["User-Agent"])))
		response.SetBody(req.headers["User-Agent"])

		return response.String()
	case strings.HasPrefix(req.path, FILES_PATH):
		file := strings.Split(req.path, "/")[2]
		directory := GetDirectoryArg()
		path := filepath.Join(directory, file)

		switch {
		case req.method == GET:
			content, err := os.ReadFile(path)

			response := NewResponse()
			response.SetProtocol(PROTOCOL)

			if err != nil {
				fmt.Println("Error reading file: ", err.Error())
				content := []byte("File not found")

				response.SetStatus(HTTP_STATUS_NOT_FOUND)
				response.SetHeader(CONTENT_TYPE_HEADER, CONTENT_TYPE_TEXT)
				response.SetHeader(CONTENT_LENGTH_HEADER, fmt.Sprintf("%d", len(content)))
				response.SetBody(string(content))

				return response.String()
			}

			response.SetStatus(HTTP_STATUS_OK)
			response.SetHeader(CONTENT_TYPE_HEADER, CONTENT_TYPE_STREAM)
			response.SetHeader(CONTENT_LENGTH_HEADER, fmt.Sprintf("%d", len(content)))
			response.SetBody(string(content))

			return response.String()
		case req.method == POST:
			file, err := os.Create(path)
			fileContents := strings.Split(request, SEPARATOR+SEPARATOR)[1]

			response := NewResponse()
			response.SetProtocol(PROTOCOL)

			if err != nil {
				fmt.Println("Error creating file: ", err.Error())
				response.SetStatus(HTTP_STATUS_INTERNAL)

				return response.String()
			}

			buf := []byte(fileContents)
			_, err = file.Write(buf)
			if err != nil {
				fmt.Println("Error writing to file: ", err.Error())
				response.SetStatus(HTTP_STATUS_INTERNAL)

				return response.String()
			}

			response.SetStatus(HTTP_STATUS_CREATED)

			return response.String()
		default:
			response := NewResponse()
			response.SetProtocol(PROTOCOL)
			response.SetStatus(HTTP_STATUS_NOT_IMPL)

			return response.String()
		}
	default:
		response := NewResponse()
		response.SetProtocol(PROTOCOL)
		response.SetStatus(HTTP_STATUS_NOT_FOUND)

		return response.String()
	}
}
