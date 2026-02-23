package main

import (
	"log"
	"net"
)

func handleConnection(conn net.Conn) {
	defer conn.Close()

	ok := "OK\n"
	buf := []byte(ok)
	_, err := conn.Write(buf)

	if err != nil {
		log.Println(err)
	}

}

func main() {
	ln, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		log.Fatalln(err)
	}

	defer ln.Close()

	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Println(err)
		} else {
			go handleConnection(conn)
		}
	}

}
