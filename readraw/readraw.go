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

	// Read the File System Structures
	// Implement your logic to read the file system structures here.
	// This would involve reading the boot sector, FAT table, and directory entries.

	// Example: Reading the first sector
	bootSector := make([]byte, 512)
	var bytesRead uint32
	err = syscall.ReadFile(
		syscall.Handle(handle),
		bootSector,
		&bytesRead,
		nil,
	)
	if err != nil {
		fmt.Println("Error reading boot sector:", err)
		return
	}

	// Parse the boot sector to get information about the file system
	// This is a simplified example and doesn't cover the entire FAT file system.

	// Traverse Directories
	// Implement your logic to traverse directories and access files here.

	// Example: Listing files in the root directory
	rootDirEntries := make([]byte, 512) // Assuming root directory is 512 bytes
	err = syscall.ReadFile(
		syscall.Handle(handle),
		rootDirEntries,
		&bytesRead,
		nil,
	)
	if err != nil {
		fmt.Println("Error reading root directory:", err)
		return
	}

	fmt.Println(rootDirEntries)

	// Parse directory entries and list files
	// This is a simplified example and doesn't cover all directory entry formats.

	// Read File Content
	// Implement your logic to read file content here.
}
