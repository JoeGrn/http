package server

import (
	"fmt"
	"net"
	"os"

	"github.com/joegrn/http/pkg/consts"
	"github.com/joegrn/http/pkg/request"
	"github.com/joegrn/http/pkg/response"
)

func Serve() {

	l, err := net.Listen("tcp", consts.URL)

	if err != nil {
		fmt.Println("Failed to bind to port: ", consts.URL)
		os.Exit(1)
	}

	defer l.Close()

	for {
		conn, err := l.Accept()
		if err != nil {
			fmt.Println("Error accepting connection: ", err.Error())
			os.Exit(1)
		}

		go HandleConnection(conn)
	}

}

func HandleConnection(conn net.Conn) {

	defer conn.Close()
	buf := make([]byte, 1024)
	for {
		size, err := conn.Read(buf)
		if err != nil {
			fmt.Println("Error reading connection: ", err.Error())
			return
		}
		req := string(buf[:size])
		parsedReq := request.ParseRequest(req)
		res := response.HandleResponse(parsedReq)

		conn.Write([]byte(res))
	}
}
