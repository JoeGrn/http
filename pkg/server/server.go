package server

import (
	"fmt"
	"net"
	"os"

	"github.com/joegrn/http/pkg/request"
	"github.com/joegrn/http/pkg/response"
)

func Serve() {

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
		req := string(buf[:size])
		parsedReq := request.ParseRequest(req)
		res := response.HandleResponse(parsedReq)

		conn.Write([]byte(res))
	}
}
