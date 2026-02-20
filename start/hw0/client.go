package main

import (
	"bufio"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		os.Exit(1)
	}
	defer conn.Close()
	r := bufio.NewReader(conn)
	line, err := r.ReadString('\n')
	if err != nil {
		os.Exit(1)
	}
	if line != "OK\n" {
		os.Exit(1)
	}
}
