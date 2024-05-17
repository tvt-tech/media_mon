//go:build !mipsle
// +build !mipsle

package utils

import (
	// "log"
	// "os"

	"os"
	"os/exec"

	"github.com/getlantern/systray"
)

func onReady() {
	// Add an item to the system tray
	systray.SetIcon(IconData) // Set the icon
	systray.SetTitle("Archer USB monitor")
	systray.SetTooltip("Archer USB monitor")
	mQuit := systray.AddMenuItem("Quit", "Quit the whole app")
	mRestart := systray.AddMenuItem("Restart", "Restart service")

	// Handle click on the restart item
	go func() {
		<-mRestart.ClickedCh
		onRestart()
	}()

	// Handle click on the quit item
	go func() {
		<-mQuit.ClickedCh
		systray.Quit()
	}()
}

func onRestart() {
	Logger.Debugf("Restarting...")

	// Get the path to the current executable
	executable, err := os.Executable()
	if err != nil {
		Logger.Error("Error getting executable path:", err)
		return
	}

	// Execute the current executable with a special argument to signal a restart
	// cmd := exec.Command(executable, "--restart")
	cmd := exec.Command(executable, "-s")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	// Start the new process
	err = cmd.Start()
	if err != nil {
		Logger.Error("Error restarting application:", err)
	} else {
		systray.Quit()
	}
}

func onExit() {
	// Clean up resources or save state if needed
	os.Exit(0)
}

func RunTray() {
	systray.Run(onReady, onExit)
}
