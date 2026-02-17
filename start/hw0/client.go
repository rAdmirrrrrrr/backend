package main

import (
	"fmt"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		panic(err)
	}

	defer conn.Close()

	buff := make([]byte, 4)
	n, _ := conn.Read(buff)
	answer := string(buff[:n])

	if answer != "OK\n" {
		fmt.Printf("answer from server is wrong")
	}
}
