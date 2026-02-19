package main

import (
	"bufio"
	"log"
	"net"
	"time"
)

const (
	addr          = "127.0.0.1:8080"
	timeout       = 10 * time.Second
	expectedReply = "OK\n"
)

func main() {
	dialer := &net.Dialer{
		Timeout: timeout,
	}

	conn, err := dialer.Dial("tcp", addr)
	if err != nil {
		log.Fatalf("failed to dial server: %v", err)
	}
	defer conn.Close()

	log.Printf("connected to %s", addr)

	reader := bufio.NewReader(conn)
	reply, err := reader.ReadString('\n')
	if err != nil {
		log.Fatalf("failed to read response: %v", err)
	}

	if reply != expectedReply {
		log.Fatalf("unexpected reply: %s", reply)
	}

	log.Println("server replied correctly")
}
