package request

import (
	"strings"

	"github.com/joegrn/http/pkg/consts"
)

func ParseHeaders(request string) map[string]string {
	headers := make(map[string]string)
	lines := strings.Split(request, "\r\n")

	for i := 1; i < len(lines); i++ {
		if lines[i] == "" {
			break
		}

		parts := strings.Split(strings.TrimSpace(lines[i]), ": ")
		headers[parts[0]] = parts[1]
	}

	return headers
}

func ParseRequest(request string) *Request {
	parts := strings.Split(request, " ")
	headers := ParseHeaders(request)
	body := strings.Split(request, consts.SEPARATOR+consts.SEPARATOR)[1]

	return &Request{
		Method:          RemoveSeparators(parts[0]),
		Path:            RemoveSeparators(parts[1]),
		ProtocolVersion: RemoveSeparators(parts[2]),
		Headers:         headers,
		Body:            body,
	}
}

func RemoveSeparators(request string) string {
	return strings.ReplaceAll(request, consts.SEPARATOR, "")
}
