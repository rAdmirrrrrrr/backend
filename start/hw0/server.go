package main

import (
	"fmt"
	"net"
)

type Server struct {
	listenAddr string
	ln         net.Listener
	quitch     chan struct{}
}

func NewServer(listenAddr string) *Server {
	return &Server{
		listenAddr: listenAddr,
		quitch:     make(chan struct{}),
	}
}

func (s *Server) Start() error {
	ln, err := net.Listen("tcp", s.listenAddr)
	if err != nil {
		return err
	}
	defer ln.Close()
	s.ln = ln

	go s.acceptLoop()

	<-s.quitch
	return nil
}

func (s *Server) acceptLoop() {
	for {
		conn, err := s.ln.Accept()
		if err != nil {
			fmt.Println("Ошибка:", err)
			continue
		}

		fmt.Println("Новое подключение: ", conn.RemoteAddr())
		_, err = conn.Write([]byte("OK\n"))
		if err != nil {
			fmt.Println("Ошибка отправки OK:", err)
			conn.Close()
			continue
		}

		go s.readLoop(conn)
	}
}

func (s *Server) readLoop(conn net.Conn) {
	defer conn.Close()
	buf := make([]byte, 2048)
	for {
		n, err := conn.Read(buf)
		if err != nil {
			fmt.Println("Соединение закрыто:", err)
			break
		}
		msg := buf[:n]
		fmt.Printf("Получено сообщение: %s\n", string(msg))

		response := fmt.Sprintf("Сервер получил: %s", string(msg))
		_, err = conn.Write([]byte(response))
		if err != nil {
			fmt.Println("Ошибка отправки ответа:", err)
			break
		}
	}
}

func main() {
	server := NewServer(":8080")
	err := server.Start()
	if err != nil {
		return
	}
}
