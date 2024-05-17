package main

import (
	"fmt"
	"syscall"

	"github.com/gonutz/w32"
)

func main() {
	// Open the device with read access
	devicePath := `\\.\PhysicalDrive1` // Change this to the appropriate device path
	handle, err := syscall.CreateFile(
		syscall.StringToUTF16Ptr(devicePath),
		w32.GENERIC_READ,
		w32.FILE_SHARE_READ|w32.FILE_SHARE_WRITE,
		nil,
		w32.OPEN_EXISTING,
		0,
		0,
	)
	if err != nil {
		fmt.Println("Error:", err)
		return
	} else {
		fmt.Println("Got handle")
	}
	defer syscall.CloseHandle(handle)

	// Read data from the device
	buffer := make([]byte, 512) // Reading a sector, adjust the buffer size as needed
	var bytesReturned uint32
	err = syscall.ReadFile(
		syscall.Handle(handle),
		buffer,
		&bytesReturned,
		nil,
	)
	if err != nil {
		fmt.Println("Error reading from device:", err)
		return
	}

	// // Process the data read from the device
	fmt.Println("Data read:", buffer[:bytesReturned])

	// You can now process the data read from the device
}
