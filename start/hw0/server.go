package main

import (
	"log"
	"net"
)

func main() {
	listenner, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}

	defer listenner.Close()

	for {
		conn, err := listenner.Accept()
		if err != nil {
			log.Println(err)
		}

		msg := []byte("OK\n")
		_, err = conn.Write(msg)
		if err != nil {
			log.Println(err)
		}

		conn.Close()
	}
}
