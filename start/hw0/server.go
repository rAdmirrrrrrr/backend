package main

import (
	"fmt"
	"net"
)

const (
	serverPort      = "8080"   
	successMessage  = "OK\n"   
	networkType     = "tcp"    
)

func main() {
	address := ":" + serverPort
	listener, err := net.Listen(networkType, address)
	if err != nil {
		fmt.Printf("Ошибка при запуске сервера на порту %s: %v\n", serverPort, err)
		return
	}
	fmt.Printf("Сервер запущен на %s:%s\n", networkType, serverPort)

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Ошибка при подключении клиента:", err)
			continue
		}

		go func(c net.Conn) {
			defer c.Close()
			_, err := c.Write([]byte(successMessage))
			if err != nil {
				fmt.Println("Ошибка при отправке данных клиенту:", err)
			} else {
				fmt.Printf("Отправлено клиенту %s: %s", c.RemoteAddr(), successMessage)
			}
		}(conn)
	}
}