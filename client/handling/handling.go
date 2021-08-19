package handling

import (
	"bufio"
	"flag"
	"fmt"
	"net"
	"strings"

	"github.com/marcusolsson/tui-go"
)

// Connects to server.
// Returns a net.Conn variable.
func ConnectToServer() net.Conn {
	var ipAddr string

	flag.StringVar(&ipAddr, "a", "127.0.0.1", "Specify ip address. Default is localhost")
	flag.Parse()

	var ip = fmt.Sprint(ipAddr, ":8642")

	// Connect to server
	conn, err := net.Dial("tcp", ip)
	if err != nil {
		panic(err)
	}

	return conn
}

func HandleMessages(connection *net.Conn, history *tui.Box, ui *tui.UI) {

	var UI = *ui

	conn := *connection
	reader := bufio.NewReader(conn)

	var isRunning = true
	for isRunning {
		message, err := reader.ReadString('\n')
		message = strings.TrimSuffix(message, "\n")
		if err != nil {
			history.Append(tui.NewHBox(
				tui.NewPadder(1, 0, tui.NewLabel(fmt.Sprintf("[%v]", "ERROR"))),
				tui.NewLabel(err.Error()),
			))
			UI.Repaint()
			isRunning = false
			return
		}
		history.Append(tui.NewHBox(
			tui.NewPadder(1, 0, tui.NewLabel(fmt.Sprintf("<%v>", "NotMe"))),
			tui.NewLabel(message),
		))
		UI.Repaint()

	}
}