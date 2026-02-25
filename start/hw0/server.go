package main

import (
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

func handleConnection(conn net.Conn) {
	defer conn.Close()

	connName := conn.RemoteAddr().String()

	_, err := conn.Write([]byte(successMessage))
	if err != nil {
		log.Printf("can not answer to client %s: %w", connName, err)
		return
	}
}

func main() {
	listener, err := net.Listen(protocol, serverAddress)
	if err != nil {
		log.Printf("can not start server: %w", err)
		return
	}

	defer listener.Close()

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("can not return the connection", err)
		}
		go handleConnection(conn)
	}
}
