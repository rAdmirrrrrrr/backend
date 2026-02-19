package main

import (
	"flag"
	"fmt"
	"log"
	"net"
)

const (
	msg = "OK\n"
)

func main() {
	serverAddr := flag.String("serverAddr", ":8080", "TCP server address")
	flag.Parse()

	ln, err := net.Listen("tcp", *serverAddr)
	if err != nil {
		log.Fatal("Internal server error", err)
	}

	defer ln.Close()

	fmt.Println("Server is running...")

	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Println("Connection error:", err)
			continue
		}
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()

	_, err := conn.Write([]byte(msg))
	if err != nil {
		log.Println("Response sending error", err)
	}
	log.Println("Response is sent", conn.RemoteAddr())
}
