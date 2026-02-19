package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
)

const (
	msg = "OK\n"
)

func main() {
	clientAddr := flag.String("clientAddr", "127.0.0.1:8080", "TCP client address")
	flag.Parse()

	conn, err := net.Dial("tcp", *clientAddr)
	if err != nil {
		log.Fatal("Error while connecting to the server", err)
	}
	defer conn.Close()

	fmt.Println("Connection established...")

	response, err := readResponse(conn)
	if err != nil {
		log.Fatal("Response reading error", err)
	}

	if strings.Compare(string(response), msg) == 0 {
		fmt.Println(string(response))
	} else {
		fmt.Println("Wrong response")
		os.Exit(1)
	}
}

func readResponse(conn net.Conn) ([]byte, error) {
	var message []byte
	buffer := make([]byte, 1)

	for {
		n, err := conn.Read(buffer)
		if err != nil {
			return message, err
		}
		if n > 0 {
			message = append(message, buffer[:n]...)
			if buffer[0] == '\n' {
				return message, err
			}
		}
	}
}
