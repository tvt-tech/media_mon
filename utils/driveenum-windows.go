//go:build windows
// +build windows

package utils

import (
	"bufio"
	"bytes"
	"os"
	"os/exec"
	"strings"
)

// Detect returns a list of file paths pointing to the root folder of
// USB storage devices connected to the system.
func Detect() ([]string, error) {
	var drives []string
	driveMap := make(map[string]bool)

	cmd := "wmic"
	args := []string{"logicaldisk", "where", "drivetype=2", "get", "deviceid"}
	out, err := exec.Command(cmd, args...).Output()

	if err != nil {
		return drives, err
	}

	s := bufio.NewScanner(bytes.NewReader(out))
	for s.Scan() {
		line := s.Text()
		if strings.Contains(line, ":") {
			rootPath := strings.TrimSpace(line) + string(os.PathSeparator)
			driveMap[rootPath] = true
		}
	}

	for k := range driveMap {
		file, err := os.Open(k)
		if err == nil {
			drives = append(drives, k)
		}
		file.Close()
	}

	return drives, nil
}

/* OPTIONAL INTERNAL USE */
// func isUSBStorage(device string) bool {
// 	deviceVerifier := "ID_USB_DRIVER=usb-storage"
// 	cmd := "udevadm"
// 	args := []string{"info", "-q", "property", "-n", device}
// 	out, err := exec.Command(cmd, args...).Output()

// 	if err != nil {
// 		Logger.Errorf("Error checking device %s: %s", device, err)
// 		return false
// 	}

// 	if strings.Contains(string(out), deviceVerifier) {
// 		return true
// 	}

// 	return false
// }

func PrintDrives() {
	if drives, err := Detect(); err == nil {
		Logger.Infof("%d USB Devices Found", len(drives))
		for i, d := range drives {
			Logger.Infof("%d: %s", i, d)
		}
	} else {
		Logger.Error(err)
	}
}
