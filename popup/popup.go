package main

import (
	"fmt"

	"github.com/ChromeTemp/Popup"
)

func main() {
	// shows a native Windows alert
	// Popup.Alert("Example Title", "Example Content")
	// shows a native Windows dialog (and handle user action)
	Popup.Dialog(
		"Dialog",
		fmt.Sprintf("Found unexpectable files on %s .\nRename them to prevent an execution?", "E:/"),
	)
}
