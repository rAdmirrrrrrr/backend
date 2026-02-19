package main

import (
	"errors"
	"log"
	"net"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

const (
	writeTimeout  = 10 * time.Second
	readTimeout   = 10 * time.Second
	ip            = "0.0.0.0"
	port          = "8080"
	protocol      = "tcp"
	expectedReply = "OK\n"
)

func main() {
	listener, err := net.Listen(protocol, net.JoinHostPort(ip, port))
	if err != nil {
		log.Println("failed to create listener, err:", err)
		os.Exit(1)
	}
	log.Printf("listening on %s\n", listener.Addr())

	var waitGroup sync.WaitGroup

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		<-stop
		log.Println("shutting down...")
		listener.Close()
	}()

	for {
		conn, err := listener.Accept()
		if err != nil {
			if errors.Is(err, net.ErrClosed) {
				break
			}
			log.Println("failed to accept connection, err:", err)
			continue
		}
		waitGroup.Add(1)
		go handleConnection(conn, &waitGroup)
	}

	waitGroup.Wait()
	log.Println("server stopped")
}

func handleConnection(conn net.Conn, waitGroup *sync.WaitGroup) {
	defer waitGroup.Done()
	defer conn.Close()

	conn.SetWriteDeadline(time.Now().Add(writeTimeout))
	conn.SetReadDeadline(time.Now().Add(readTimeout))

	if _, err := conn.Write([]byte(expectedReply)); err != nil {
		log.Println("failed to write to connection:", err)
		return
	}
}
