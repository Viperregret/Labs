package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		fmt.Println("Ошибка подключения к серверу:", err)
		return
	}
	defer conn.Close()

	// Ввод сообщения от пользователя
	fmt.Print("Введите сообщение для отправки серверу: ")
	message, _ := bufio.NewReader(os.Stdin).ReadString('\n')

	// Отправка сообщения серверу
	fmt.Fprintf(conn, message)

	// Получение подтверждения от сервера
	response, _ := bufio.NewReader(conn).ReadString('\n')
	fmt.Printf("Ответ от сервера: %s", response)
}
