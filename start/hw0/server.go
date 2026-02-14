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

const writeTimeout = 10 * time.Second
const ip = "0.0.0.0"
const port = "8080"
const protocol = "tcp"

type connectionsSet struct {
	mutex       sync.Mutex
	connections map[net.Conn]struct{}
}

func (s *connectionsSet) add(conn net.Conn) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	if s.connections == nil {
		s.connections = make(map[net.Conn]struct{})
	}
	s.connections[conn] = struct{}{}
}

func (s *connectionsSet) remove(conn net.Conn) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	delete(s.connections, conn)
}

func (s *connectionsSet) closeAll() {
	s.mutex.Lock()
	list := make([]net.Conn, 0, len(s.connections))
	for c := range s.connections {
		list = append(list, c)
	}
	s.mutex.Unlock()
	for _, c := range list {
		c.Close()
	}
}

func main() {
	listener, err := net.Listen(protocol, net.JoinHostPort(ip, port))
	if err != nil {
		log.Println("failed to create listener, err:", err)
		os.Exit(1)
	}
	log.Printf("listening on %s\n", listener.Addr())

	var waitGroup sync.WaitGroup
	connections := &connectionsSet{}

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		<-stop
		log.Println("shutting down...")
		listener.Close()
		connections.closeAll()
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
		connections.add(conn)
		waitGroup.Add(1)
		go handleConnection(conn, &waitGroup, connections)
	}

	waitGroup.Wait()
	log.Println("server stopped")
}

func handleConnection(conn net.Conn, waitGroup *sync.WaitGroup, connections *connectionsSet) {
	defer waitGroup.Done()
	defer connections.remove(conn)
	defer conn.Close()

	conn.SetWriteDeadline(time.Now().Add(writeTimeout))
	if _, err := conn.Write([]byte("OK\n")); err != nil {
		return
	}
}
