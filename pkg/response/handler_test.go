package response_test

import (
	"fmt"
	"testing"

	"github.com/joegrn/http/pkg/consts"
	"github.com/joegrn/http/pkg/request"
	"github.com/joegrn/http/pkg/response"
)

func mockCompressor(body string) string {
	return body
}

func TestHandleRoot(t *testing.T) {
	resp := response.HandleRoot()
	expectedString := "HTTP/1.1 200 OK\r\n\r\n"
	if resp != expectedString {
		t.Errorf("Expected string %s, got %s", expectedString, resp)
	}
}

func TestHandleEcho(t *testing.T) {
	req := &request.Request{
		Path:    consts.ECHO_PATH + "Hello",
		Headers: map[string]string{},
	}
	resp := response.HandleEcho(req, mockCompressor)
	expectedString := "HTTP/1.1 200 OK\r\nContent-Type: text/plain\r\nContent-Length: 5\r\n\r\nHello"
	if resp != expectedString {
		t.Errorf("Expected string %s, got %s", expectedString, resp)
	}
}

func TestHandleEchoWithGzip(t *testing.T) {
	req := &request.Request{
		Path: consts.ECHO_PATH + "Hello",
		Headers: map[string]string{
			consts.ACCEPT_ENCODING_HEADER: consts.ENCODING_TYPE_GZIP,
		},
	}

	body := mockCompressor("Hello")
	resp := response.HandleEcho(req, mockCompressor)
	expectedString := "HTTP/1.1 200 OK\r\nContent-Type: text/plain\r\nContent-Encoding: gzip\r\nContent-Length: %d\r\n\r\n%s"
	expectedString = fmt.Sprintf(expectedString, len(body), body)
	if resp != expectedString {
		t.Errorf("Expected string %s, got %s", expectedString, resp)
	}
}

func TestHandleUserAgent(t *testing.T) {
	req := &request.Request{
		Headers: map[string]string{
			"User-Agent": "Go-http-client/1.1",
		},
	}
	resp := response.HandleUserAgent(req)
	expectedString := "HTTP/1.1 200 OK\r\nContent-Type: text/plain\r\nContent-Length: 18\r\n\r\nGo-http-client/1.1"
	if resp != expectedString {
		t.Errorf("Expected string %s, got %s", expectedString, resp)
	}
}

func TestHandleNotFound(t *testing.T) {
	resp := response.HandleNotFound()
	expectedString := "HTTP/1.1 404 Not Found\r\n\r\n"
	if resp != expectedString {
		t.Errorf("Expected string %s, got %s", expectedString, resp)
	}
}

func TestNotImplemented(t *testing.T) {
	req := &request.Request{
		Method: "PATCH",
		Path:   "/files/test.txt",
	}
	resp := response.HandleFiles(req)
	expectedString := "HTTP/1.1 501 Not Implemented\r\n\r\n"
	if resp != expectedString {
		t.Errorf("Expected string %s, got %s", expectedString, resp)
	}
}
