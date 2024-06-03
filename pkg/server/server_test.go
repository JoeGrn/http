package server_test

import (
	"net"
	"testing"

	"github.com/joegrn/http/pkg/server"
)

func TestHandleConnection(t *testing.T) {
	s, client := net.Pipe()

	go func() {
		server.HandleConnection(s)
	}()

	request := "GET / HTTP/1.1\r\nHost: localhost\r\n\r\n"
	client.Write([]byte(request))

	buf := make([]byte, 1024)
	n, err := client.Read(buf)
	if err != nil {
		t.Fatal("Error reading from client: ", err)
	}

	actualResponse := string(buf[:n])
	expectedResponse := "HTTP/1.1 200 OK\r\n\r\n"
	if actualResponse != expectedResponse {
		t.Errorf("Expected '%s' but got '%s'", expectedResponse, actualResponse)
	}

	client.Close()
}
