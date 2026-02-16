package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:8080")
	if err != nil {
		fmt.Println("Ошибка подключения:", err)
		os.Exit(1)
	}
	defer conn.Close()

	reply := make([]byte, 1024)
	n, err := conn.Read(reply)
	if err != nil {
		fmt.Println("Ошибка чтения приветствия:", err)
		os.Exit(1)
	}
	fmt.Printf("Сервер: %s", string(reply[:n]))

	testMessage := "Привет"
	_, err = conn.Write([]byte(testMessage + "\n"))
	if err != nil {
		fmt.Println("Ошибка отправки:", err)
		os.Exit(1)
	}
	fmt.Printf("Отправлено: %s\n", testMessage)

	n, err = conn.Read(reply)
	if err != nil {
		fmt.Println("Ошибка получения ответа:", err)
		os.Exit(1)
	}
	fmt.Printf("Получен ответ: %s\n", string(reply[:n]))

	fmt.Println("Клиент успешно выполнился")
}
