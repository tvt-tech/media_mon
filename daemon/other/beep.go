// Copyright 2012 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build windows

package main

import (
	"bufio"
	"bytes"
	"log"
	"os"
	"os/exec"
	"strings"
	"syscall"
	"time"
)

// BUG(brainman): MessageBeep Windows api is broken on Windows 7,
// so this example does not beep when runs as service on Windows 7.

var connected_drives []string

var (
	beepFunc = syscall.MustLoadDLL("user32.dll").MustFindProc("MessageBeep")
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

func upd_drives() {
	if drives, err := Detect(); err != nil {
		upd_drives()
	} else {
		connected_drives = drives
	}
}

func init() {
	logFile, err := os.OpenFile("C:\\beep.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatalf("Failed to open log file: %v", err)
	}
	log.SetOutput(logFile)

	upd_drives()

}

func isInList(search string, list []string) bool {
	for _, item := range list {
		if item == search {
			return true // If found, return true
		}
	}
	return false // If not found, return false
}

func beep() {
	// beepFunc.Call(0xffffffff)
	prev_drives := make([]string, len(connected_drives))
	copy(prev_drives, connected_drives)

	upd_drives()
	if connected_drives != nil {
		for _, nd := range connected_drives {
			if !isInList(nd, prev_drives) {
				log.Println("New connected: %s", nd)
			}
		}
	}
}

func main() {
	for {
		time.Sleep(1 * time.Second)
		beep()
	}

}
