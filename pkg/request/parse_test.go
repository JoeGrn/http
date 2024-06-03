package request_test

import (
	"reflect"
	"testing"

	"github.com/joegrn/http/pkg/request"
)

func TestParseRequest(t *testing.T) {
	testCases := []struct {
		desc     string
		request  string
		expected *request.Request
	}{
		{
			desc:    "GET request no headers or body",
			request: "GET / HTTP/1.1\r\n\r\n",
			expected: &request.Request{
				Method:          "GET",
				Path:            "/",
				ProtocolVersion: "HTTP/1.1",
				Headers:         map[string]string{},
				Body:            "",
			},
		},
		{
			desc:    "POST request with headers and body",
			request: "POST /echo HTTP/1.1\r\n Content-Type: application/json\r\nContent-Length: 13\r\n\r\n{\"key\":\"val\"}",
			expected: &request.Request{
				Method:          "POST",
				Path:            "/echo",
				ProtocolVersion: "HTTP/1.1",
				Headers: map[string]string{
					"Content-Type":   "application/json",
					"Content-Length": "13",
				},
				Body: "{\"key\":\"val\"}",
			},
		},
	}

	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			result := request.ParseRequest(tC.request)
			if !reflect.DeepEqual(result, tC.expected) {
				t.Errorf("expected %+v, got %+v", tC.expected, result)
			}
		})
	}
}
