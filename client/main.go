package main

import (
	"fmt"
	"github.com/marcusolsson/tui-go"
	"clat/client/handling"
)



func main() {

	conn := handling.ConnectToServer()


	// MESSAGE HISTORY
	history := tui.NewVBox(
		tui.NewLabel("No Messages RightNow"),
	)
	// create history scroll area for auto scrolling
	historyScroll := tui.NewScrollArea(history)
	historyScroll.SetAutoscrollToBottom(true)
	// create history box (to copy the tutorial)
	historyBox := tui.NewVBox(historyScroll)
	historyBox.SetBorder(true)
	historyBox.SetTitle("Messages")

	// INPUT FIELD
	input := tui.NewEntry()
	input.SetFocused(true)
	input.SetSizePolicy(tui.Expanding, tui.Maximum)

	inputBox := tui.NewHBox(input)
	inputBox.SetBorder(true)
	inputBox.SetTitle("Input")
	inputBox.SetSizePolicy(tui.Expanding, tui.Maximum)

	// Input submit
	input.OnSubmit(func(entry *tui.Entry) {
		fmt.Fprint(conn, entry.Text(), "\n")
		history.Append(tui.NewHBox(
			tui.NewPadder(1, 0, tui.NewLabel(fmt.Sprintf("<%v>", "Me"))),
			tui.NewLabel(entry.Text()),
		))
		input.SetText("")
	})


	// BASE WIDGET
	box := tui.NewVBox(
		historyBox,
		inputBox,
	)


	// create root widget
	ui, err := tui.New(box)
	if err != nil {
		panic(err)
	}

	ui.SetKeybinding("esc", func () { ui.Quit() })  // Quits on pressing escape key

	// Handle messages
	go handling.HandleMessages(&conn, history, &ui)

	// run root widget
	if err := ui.Run(); err != nil {
		panic(err)
	}
}