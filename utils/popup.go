package utils

import (
	"fmt"

	"github.com/ChromeTemp/Popup"
)

func PopupDoRename(path string, files []string) bool {
	// shows a native Windows alert
	// Popup.Alert("Example Title", "Example Content")
	// shows a native Windows dialog (and handle user action)
	// Example: user will press Ok
	// fmt.Printf("Pressed Ok? %d", res)
	// logs: Pressed Ok? true
	return Popup.Dialog(
		"Warning",
		fmt.Sprintf("Found unexpected files on %s .\nRename them to prevent an execution?", path),
	)
}
