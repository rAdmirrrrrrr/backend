package main

import (
	"fmt"
	"io"
	"net"
	"os"
)

const (
	serverAddress   = "localhost:8080" 
	expectedMessage = "OK\n"           
	bufferSize      = 1024             
)

func main() {
	conn, err := net.Dial("tcp", serverAddress)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Ошибка подключения к серверу %s: %v\n", serverAddress, err)
		return
	}
	defer conn.Close()

	buf := make([]byte, bufferSize)
	n, err := conn.Read(buf)
	if err != nil && err != io.EOF {
		fmt.Fprintf(os.Stderr, "Ошибка чтения данных: %v\n", err)
		return
	}

	data := string(buf[:n])

	if data != expectedMessage {
		fmt.Fprintf(os.Stderr, "Получен неожиданный ответ: %s\n", data)
		return
	}

	fmt.Println("Ответ сервера получен успешно:", data)
}