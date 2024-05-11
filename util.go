package main

import (
	"compress/gzip"
	"os"
	"strings"
)

func ParseHeaders(request string) map[string]string {
	headers := make(map[string]string)
	lines := strings.Split(request, "\r\n")

	for i := 1; i < len(lines); i++ {
		if lines[i] == "" {
			break
		}

		parts := strings.Split(lines[i], ": ")
		headers[parts[0]] = parts[1]
	}

	return headers
}

func ParseRequest(request string) *Request {
	parts := strings.Split(request, " ")
	headers := ParseHeaders(request)

	return &Request{
		method:          parts[0],
		path:            parts[1],
		headers:         headers,
		protocolVersion: parts[2],
	}
}

func GetDirectoryArg() string {
	args := os.Args
	for i := 1; i < len(args)-1; i++ {
		if args[i] == "--directory" {
			return args[i+1]
		}
	}

	return ""
}

func GzipCompress(data string) string {
	var b strings.Builder
	w := gzip.NewWriter(&b)
	w.Write([]byte(data))
	w.Close()

	return b.String()
}
