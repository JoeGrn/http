package response

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/joegrn/http/pkg/consts"
	"github.com/joegrn/http/pkg/request"
	"github.com/joegrn/http/pkg/util"
)

func HandleRoot() string {
	response := NewResponse()
	response.SetProtocol(consts.PROTOCOL)
	response.SetStatus(consts.HTTP_STATUS_OK)
	return response.String()
}

func HandleEcho(req *request.Request, compressor util.Compressor) string {
	body := strings.ReplaceAll(req.Path, consts.ECHO_PATH, "")

	response := NewResponse()
	response.SetProtocol(consts.PROTOCOL)
	response.SetStatus(consts.HTTP_STATUS_OK)
	response.SetHeader(consts.CONTENT_TYPE_HEADER, consts.CONTENT_TYPE_TEXT)

	if strings.Contains(req.Headers[consts.ACCEPT_ENCODING_HEADER], consts.ENCODING_TYPE_GZIP) {
		response.SetHeader(consts.CONTENT_ENCODING_HEADER, consts.ENCODING_TYPE_GZIP)
		body = compressor(body)

	}

	response.SetHeader(consts.CONTENT_LENGTH_HEADER, fmt.Sprintf("%d", len(body)))
	response.SetBody(body)

	return response.String()
}

func HandleUserAgent(req *request.Request) string {
	response := NewResponse()
	response.SetProtocol(consts.PROTOCOL)
	response.SetStatus(consts.HTTP_STATUS_OK)
	response.SetHeader(consts.CONTENT_TYPE_HEADER, consts.CONTENT_TYPE_TEXT)
	response.SetHeader(consts.CONTENT_LENGTH_HEADER, fmt.Sprintf("%d", len(req.Headers["User-Agent"])))
	response.SetBody(req.Headers["User-Agent"])

	return response.String()
}

func HandleNotFound() string {
	response := NewResponse()
	response.SetProtocol(consts.PROTOCOL)
	response.SetStatus(consts.HTTP_STATUS_NOT_FOUND)

	return response.String()
}

func HandleFiles(req *request.Request) string {
	file := strings.Split(req.Path, "/")[2]
	directory := util.GetDirectoryArg()
	path := filepath.Join(directory, file)

	switch {
	case req.Method == consts.GET:
		content, err := os.ReadFile(path)

		response := NewResponse()
		response.SetProtocol(consts.PROTOCOL)

		if err != nil {
			fmt.Println("Error reading file: ", err.Error())
			content := []byte("File not found")

			response.SetStatus(consts.HTTP_STATUS_NOT_FOUND)
			response.SetHeader(consts.CONTENT_TYPE_HEADER, consts.CONTENT_TYPE_TEXT)
			response.SetHeader(consts.CONTENT_LENGTH_HEADER, fmt.Sprintf("%d", len(content)))
			response.SetBody(string(content))

			return response.String()
		}

		response.SetStatus(consts.HTTP_STATUS_OK)
		response.SetHeader(consts.CONTENT_TYPE_HEADER, consts.CONTENT_TYPE_STREAM)
		response.SetHeader(consts.CONTENT_LENGTH_HEADER, fmt.Sprintf("%d", len(content)))
		response.SetBody(string(content))

		return response.String()
	case req.Method == consts.POST:
		file, err := os.Create(path)
		fileContents := req.Body

		response := NewResponse()
		response.SetProtocol(consts.PROTOCOL)

		if err != nil {
			fmt.Println("Error creating file: ", err.Error())
			response.SetStatus(consts.HTTP_STATUS_INTERNAL)

			return response.String()
		}

		buf := []byte(fileContents)
		_, err = file.Write(buf)
		if err != nil {
			fmt.Println("Error writing to file: ", err.Error())
			response.SetStatus(consts.HTTP_STATUS_INTERNAL)

			return response.String()
		}

		response.SetStatus(consts.HTTP_STATUS_CREATED)

		return response.String()
	default:
		response := NewResponse()
		response.SetProtocol(consts.PROTOCOL)
		response.SetStatus(consts.HTTP_STATUS_NOT_IMPL)

		return response.String()
	}
}

func HandleResponse(req *request.Request) string {
	switch {
	case req.Path == consts.ROOT_PATH:
		return HandleRoot()
	case strings.HasPrefix(req.Path, consts.ECHO_PATH):
		return HandleEcho(req, util.GzipCompress)
	case strings.HasPrefix(req.Path, consts.USER_AGENT_PATH):
		return HandleUserAgent(req)
	case strings.HasPrefix(req.Path, consts.FILES_PATH):
		return HandleFiles(req)
	default:
		return HandleNotFound()
	}
}
