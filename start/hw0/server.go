package main

import (
	"io"
	"log"
	"net"
)

func main() {
	server, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}
	defer server.Close()
	for {
		conn, err := server.Accept()
		if err != nil {
			continue
		}
		go func(c net.Conn) {
			_, _ = io.WriteString(c, "OK\n")
			defer c.Close()
		}(conn)
	}
}
