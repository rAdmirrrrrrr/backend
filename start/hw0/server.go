package main

import (
	"log"
	"net"
)


func main() {
	listener, err := net.Listen("tcp", "localhost:8080");
	if err != nil {
		log.Fatal(err)
	}
	defer listener.Close()

	log.Println("Server started on :8080")

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println("accept error:", err)
			continue
		}

		go connection(conn)
	}
}

func connection(conn net.Conn) {
	defer conn.Close()

	_, err := conn.Write([]byte("OK\n"))
	if err != nil {
		return
	}
}