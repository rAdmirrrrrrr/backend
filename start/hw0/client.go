package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

const OK = "OK\n"

func main() {
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		log.Fatal(err)
	}

	defer conn.Close()

	msg, err := bufio.NewReader(conn).ReadString('\n')
	if err != nil {
		log.Panic(err)
	}

	if string(msg) != OK {
		log.Println(fmt.Errorf("response: %s", string(msg)))
	}
}
