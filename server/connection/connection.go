package connection

import (
	"bufio"
	"fmt"
	"net"
)

var List []net.Conn

func HandleConn(connection *net.Conn) {
	conn := *connection

	// Create reader to read from connection 
	var reader = bufio.NewReader(conn)


	// Loop
	var connected = true
	for connected {
		message, err := reader.ReadString('\n')
		if err != nil {
			fmt.Fprintf(conn, "Error: %v", err)
			connected = true
			return
		}
		SendToAll(message, connection)
	}
}

func SendToAll(message string, sender *net.Conn) {
	for _, conn := range List {
		if conn != *sender{
			fmt.Fprint(conn, message)
		}
	}
}