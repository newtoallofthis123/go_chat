package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
)

func sendMessage(conn net.Conn, msg, name string) {
	conn.Write([]byte(name + ": " + msg))
}

func read(msg string) string {
	fmt.Print(msg)

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	message := scanner.Text()

	return message
}

func main() {
	fmt.Println("Connecting to server...")

	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		log.Fatal(err)
	}

	defer conn.Close()

	name := read("Enter your name: ")

	log.Default().Println("Connected to server")

	for {
		message := read("Enter your message: ")

		sendMessage(conn, message, name)
		if message == "exit" {
			break
		}
	}
}
