package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

const (
	HOST_HOST = " 127.0.0.1"
	HOST_PORT = "8081"
	HOST_TYPE = "tcp"
)

func main() {

	// Подключаемся к сокету
	conn, _ := net.Dial(HOST_TYPE, HOST_HOST+":"+HOST_PORT)
	for {
		// Чтение входных данных от stdin
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Text to send: ")
		text, _ := reader.ReadString('\n')
		// Отправляем в socket
		fmt.Fprintf(conn, text+"\n")
		// Прослушиваем ответ
		message, _ := bufio.NewReader(conn).ReadString('\n')
		fmt.Print("Message from server: " + message)
	}
}
