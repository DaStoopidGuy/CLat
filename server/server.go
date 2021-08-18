package main

import (
	"clat/server/connection"
	"fmt"
	"net"
)


func main() {
	// Listening on port 8642
	server, err := net.Listen("tcp", ":8642")
	if err != nil {
		panic(err)
	}
	fmt.Println("Listening on Port:[8642]")

	// Infinite loop
	var isRunning = true
	for isRunning {
		// Accept connection
		conn, err := server.Accept()
		if err != nil {
			fmt.Println(err)
			isRunning = false
		}

		connection.List = append(connection.List, conn)
		go connection.HandleConn(&conn)
		fmt.Println("New Connection: ", conn.RemoteAddr())

		// // creating reader to read from the connection
		// reader := bufio.NewReader(conn)
		// // getting message
		// message, _ := reader.ReadString('\n')
		// fmt.Println("Received: ", message)
	}
}