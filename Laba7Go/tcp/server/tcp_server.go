package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

func handleConnection(conn net.Conn, wg *sync.WaitGroup) {
	defer conn.Close()
	defer wg.Done()

	message, _ := bufio.NewReader(conn).ReadString('\n')
	fmt.Printf("Получено сообщение от клиента: %s", message)
	conn.Write([]byte("Сообщение получено\n"))
}

func main() {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println("Ошибка при запуске сервера:", err)
		os.Exit(1)
	}
	defer listener.Close()

	var wg sync.WaitGroup

	// Канал для перехвата системного сигнала для завершения работы
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	go func() {
		<-quit
		fmt.Println("Завершение работы сервера...")
		listener.Close()
	}()

	for {
		conn, err := listener.Accept()
		if err != nil {
			select {
			case <-quit:
				wg.Wait()
				fmt.Println("Все соединения завершены.")
				return
			default:
				fmt.Println("Ошибка соединения:", err)
			}
			continue
		}
		wg.Add(1)
		go handleConnection(conn, &wg)
	}
}
