package main

import (
	"bufio"
	"log"
	"net"
	"os"
	"time"
)

const addr = "127.0.0.1:8080"
const timeout = 10 * time.Second

func main() {
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		log.Fatalf("failed to dial server: %v", err)
	}
	defer conn.Close()

	log.Printf("connected to %s", addr)

	reader := bufio.NewReader(conn)
	reply, err := reader.ReadString('\n')
	if err != nil {
		os.Exit(1)
	}

	if reply != "OK\n" {
		os.Exit(1)
	}
}
