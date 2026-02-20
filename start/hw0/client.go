package main

import (
	"net"
	"log"
	"io"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	data, err := io.ReadAll(conn)
	if err != nil {
		log.Fatal(err)
	}
	if string(data) != "OK\n" {
		log.Fatal("unexpected response")
	}

}
