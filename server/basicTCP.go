package main

import (
	"fmt"
	"net"
)

func main() {
	listener, err := net.Listen("tcp", "0.0.0.0:8000")
	if err != nil {
		return
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			return
		}
		go handle(conn)
	}
}

func handle(conn net.Conn) {
	for {
		buf := make([]byte, 512)
		_, err := conn.Read(buf)
		if err != nil {
			return
		}
		fmt.Println("received:", string(buf))
	}
}
