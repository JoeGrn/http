package response_test

import (
	"testing"

	"github.com/joegrn/http/pkg/consts"
	"github.com/joegrn/http/pkg/response"
)

func TestNewResponse(t *testing.T) {
	resp := response.NewResponse()
	if resp == nil {
		t.Error("Expected non-nil Response")
	}
	if len(resp.Headers) != 0 {
		t.Errorf("Expected Headers to be empty, got %d", len(resp.Headers))
	}
}

func TestSetStatus(t *testing.T) {
	resp := response.NewResponse()
	status := "200 OK"
	resp.SetStatus(status)
	if resp.Status != status {
		t.Errorf("Expected status %s, got %s", status, resp.Status)
	}
}

func TestSetProtocol(t *testing.T) {
	resp := response.NewResponse()
	protocol := "HTTP/1.1"
	resp.SetProtocol(protocol)
	if resp.Protocol != protocol {
		t.Errorf("Expected protocol %s, got %s", protocol, resp.Protocol)
	}
}

func TestSetHeader(t *testing.T) {
	resp := response.NewResponse()
	key := "Content-Type"
	value := "application/json"
	resp.SetHeader(key, value)
	if val, ok := resp.Headers[key]; !ok || val != value {
		t.Errorf("Expected header %s to be %s, got %s", key, value, val)
	}
}

func TestSetBody(t *testing.T) {
	resp := response.NewResponse()
	body := "{\"key\": \"value\"}"
	resp.SetBody(body)
	if resp.Body != body {
		t.Errorf("Expected body %s, got %s", body, resp.Body)
	}
}

func TestString(t *testing.T) {
	resp := response.NewResponse()
	protocol := "HTTP/1.1"
	status := "200 OK"
	key := "Content-Type"
	value := "application/json"
	body := "{\"key\": \"value\"}"

	resp.SetProtocol(protocol)
	resp.SetStatus(status)
	resp.SetHeader(key, value)
	resp.SetBody(body)

	expectedString := protocol + status + consts.SEPARATOR +
		key + ": " + value + consts.SEPARATOR +
		consts.SEPARATOR + body

	if resp.String() != expectedString {
		t.Errorf("Expected string %s, got %s", expectedString, resp.String())
	}
}
