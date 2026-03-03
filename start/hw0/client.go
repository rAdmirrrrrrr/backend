package main

import (
	"io"
	"log"
	"net"
)

const (
	host           = "localhost"
	port           = "8080"
	protocol       = "tcp"
	serverAddress  = host + ":" + port
	successMessage = "OK\n"
)

func main() {
	conn, err := net.Dial(protocol, serverAddress)
	if err != nil {
		log.Printf("can not connect to server: %w", err)
		return
	}

	defer conn.Close()

	data, err := io.ReadAll(conn)
	if err != nil {
		log.Printf("can not read from server: %w", err)
		return
	}

	answer := string(data)

	if answer != successMessage {
		log.Printf("answer from server is wrong")
	}
}
