package main

import (
	"bufio"
	"flag"
	"fmt"
	"net"
	"os"
)


func main() {
	var ipAddr string

	flag.StringVar(&ipAddr, "a", "127.0.0.1", "Specify ip address. Default is localhost")
	flag.Parse()

	var ip = fmt.Sprint(ipAddr, ":8642")

	// Connect to server
	conn, err := net.Dial("tcp", ip)
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