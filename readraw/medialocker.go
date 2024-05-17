package utils

import (
	"fmt"
	"syscall"
	"unsafe"
)

const (
	GENERIC_READ  = 0x80000000
	GENERIC_WRITE = 0x40000000
	OPEN_EXISTING = 3
)

var (
	kernel32, _   = syscall.LoadLibrary("kernel32.dll")
	createFile, _ = syscall.GetProcAddress(kernel32, "CreateFileW")
	// deviceIoControl, _    = syscall.GetProcAddress(kernel32, "DeviceIoControl")
	closeHandle, _ = syscall.GetProcAddress(kernel32, "CloseHandle")
	// storageLockMedia, _   = syscall.GetProcAddress(kernel32, "DeviceIoControl")
	// storageUnlockMedia, _ = syscall.GetProcAddress(kernel32, "DeviceIoControl")
	storageMedia, _ = syscall.GetProcAddress(kernel32, "DeviceIoControl")
)

func lockMedia(handle uintptr) {
	var bytesReturned uint32
	ret, _, err := syscall.SyscallN(
		// storageLockMedia,
		uintptr(storageMedia),
		6,
		handle,
		// 0x2D6D0C, // IOCTL_STORAGE_MEDIA_REMOVAL
		uintptr(0x2D6D0C), // IOCTL_STORAGE_MEDIA_REMOVAL
		0,
		0,
		uintptr(unsafe.Pointer(&bytesReturned)),
		0,
	)
	if ret == 0 {
		fmt.Printf("Error: %v\n", err)
	}
}

func unlockMedia(handle uintptr) {
	fmt.Println(handle)

	var bytesReturned uint32
	ret, _, err := syscall.SyscallN(
		// uintptr(storageUnlockMedia),
		uintptr(storageMedia),
		4,
		handle,
		uintptr(0x2D6D10), // IOCTL_STORAGE_MEDIA_REMOVAL
		0,
		0,
		uintptr(unsafe.Pointer(&bytesReturned)),
		0,
	)
	if ret == 0 {
		fmt.Printf("Error: %v\n", err)
	}
}

func Lock(drive string) {
	// devicePath := `\\.\E:` // Change E: to the appropriate drive letter
	devicePath := `\\.\` + drive
	handle, _, _ := syscall.SyscallN(
		uintptr(createFile),
		7,
		uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(devicePath))),
		uintptr(GENERIC_READ|GENERIC_WRITE),
		0,
		0,
		uintptr(OPEN_EXISTING),
		0,
	)
	if handle == 0 {
		fmt.Println("Error opening device")
		return
	}
	defer syscall.SyscallN(uintptr(closeHandle), 1, handle, 0, 0)

	lockMedia(handle)
	fmt.Println("USB drive locked successfully")

	// Unlock the media when you're done
	// unlockMedia(handle)
}

func Unlock(drive string) {
	// devicePath := `\\.\E:` // Change E: to the appropriate drive letter
	devicePath := `\\.\` + drive
	handle, _, _ := syscall.SyscallN(
		uintptr(createFile),
		7,
		uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(devicePath))),
		uintptr(GENERIC_READ|GENERIC_WRITE),
		0,
		0,
		uintptr(OPEN_EXISTING),
		0,
	)
	if handle == 0 {
		fmt.Println("Error opening device")
		return
	}
	defer syscall.SyscallN(uintptr(closeHandle), 1, handle, 0, 0)

	// lockMedia(handle)
	// fmt.Println("USB drive locked successfully")

	// Unlock the media when you're done
	unlockMedia(handle)
	fmt.Println("USB drive unlocked successfully")
}
