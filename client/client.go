package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)


func main() {
	// Connect to server
	conn, err := net.Dial("tcp", ":8642")
	if err != nil {
		panic(err)
	}

	var isRunning = true
	go handleMessages(&conn, &isRunning)

	reader := bufio.NewReader(os.Stdin)
	for isRunning {
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error: ", err)
			isRunning = false
		}

		// Send message
		fmt.Fprint(conn, input)
	}
}

func handleMessages(connection *net.Conn, isRunning *bool) {

	conn := *connection
	reader := bufio.NewReader(conn)

	for *isRunning {
		message, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error: ", err)
			*isRunning = false
			return
		}
		fmt.Print(message)
	}
}