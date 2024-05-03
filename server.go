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
		return PROTOCOL + HTTP_STATUS_OK + SEPARATOR + SEPARATOR
	case strings.HasPrefix(req.path, ECHO_PATH):
		body := strings.ReplaceAll(req.path, ECHO_PATH, "")
		return PROTOCOL + HTTP_STATUS_OK + SEPARATOR + CONTENT_TYPE_HEADER + CONTENT_TYPE_TEXT + SEPARATOR + CONTENT_LENGTH_HEADER + fmt.Sprintf("%d", len(body)) + SEPARATOR + SEPARATOR + body
	case strings.HasPrefix(req.path, USER_AGENT_PATH):
		return PROTOCOL + HTTP_STATUS_OK + SEPARATOR + CONTENT_TYPE_HEADER + CONTENT_TYPE_TEXT + SEPARATOR + CONTENT_LENGTH_HEADER + fmt.Sprintf("%d", len(req.headers["User-Agent"])) + SEPARATOR + SEPARATOR + req.headers["User-Agent"]
	case strings.HasPrefix(req.path, FILES_PATH):
		file := strings.Split(req.path, "/")[2]
		directory := GetDirectoryArg()
		path := filepath.Join(directory, file)

		switch {
		case req.method == GET:
			content, err := os.ReadFile(path)

			if err != nil {
				fmt.Println("Error reading file: ", err.Error())
				content := []byte("File not found")
				return PROTOCOL + HTTP_STATUS_NOT_FOUND + SEPARATOR + CONTENT_TYPE_HEADER + CONTENT_TYPE_TEXT + SEPARATOR + CONTENT_LENGTH_HEADER + fmt.Sprintf("%d", len(content)) + SEPARATOR + SEPARATOR + string(content)
			}

			return PROTOCOL + HTTP_STATUS_OK + SEPARATOR + CONTENT_TYPE_HEADER + CONTENT_TYPE_STREAM + SEPARATOR + CONTENT_LENGTH_HEADER + fmt.Sprintf("%d", len(content)) + SEPARATOR + SEPARATOR + string(content)
		case req.method == POST:
			file, err := os.Create(path)
			fileContents := strings.Split(request, SEPARATOR+SEPARATOR)[1]

			if err != nil {
				fmt.Println("Error creating file: ", err.Error())
				return PROTOCOL + HTTP_STATUS_INTERNAL + SEPARATOR + SEPARATOR
			}

			buf := []byte(fileContents)
			_, err = file.Write(buf)
			if err != nil {
				fmt.Println("Error writing to file: ", err.Error())
				return PROTOCOL + HTTP_STATUS_INTERNAL + SEPARATOR + SEPARATOR
			}
			return PROTOCOL + HTTP_STATUS_CREATED + SEPARATOR + SEPARATOR
		default:
			return PROTOCOL + HTTP_STATUS_NOT_IMPL + SEPARATOR + SEPARATOR
		}
	default:
		return PROTOCOL + HTTP_STATUS_NOT_FOUND + SEPARATOR + SEPARATOR
	}
}
