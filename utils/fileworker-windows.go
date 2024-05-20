//go:build windows
// +build windows

package utils

import (
	"bufio"
	"fmt"
	"os"
	"path"
	"path/filepath"
	"regexp"
	"strings"
)

var allowedExtensions = []string{
	"",
	".zip",
	".bak",
	".bmp",
	".a7p",
	".jpg",
	".upg",
	".mp4",
	".tar",
	".upg",
	".tar.gz",
	".mp4",
	".jpeg",
	".png",
	".svg",
	".txt",
	".md",
}

func checkFileExist(filePath string) bool {
	// Check if file exists
	if _, err := os.Stat(filePath); err == nil {
		Logger.Debug("File exists. Path:", filePath)
		return true
	} else if os.IsNotExist(err) {
		Logger.Debug("File does not exist.")
	} else {
		Logger.Debug("Error occurred while checking file status:", err)
	}
	return false
}

func readTextFile(filePath string) (string, error) {
	// Open the file
	file, err := os.Open(filePath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	// Create a scanner to read the file line by line
	scanner := bufio.NewScanner(file)

	// Read the file line by line and append to content
	var content string
	for scanner.Scan() {
		content += scanner.Text() + "\n"
	}

	// Check for any errors during scanning
	if err := scanner.Err(); err != nil {
		return "", err
	}

	return content, nil
}

func findUUID(input string) (bool, string) {
	// Regular expression for UUID pattern
	uuidPattern := `^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}[\n$]?`

	// Compile the regex
	regex := regexp.MustCompile(uuidPattern)
	match := regex.MatchString(input)
	// Find the first UUID in the input string
	uuid := regex.FindString(input)
	return match, uuid
}

func isHasName(input string) bool {
	// Regular expression for UUID pattern
	pattern := `name_device="Archer`

	// Compile the regex
	regex := regexp.MustCompile(pattern)
	match := regex.MatchString(input)
	// Find the first UUID in the input string
	return match
}

func findUUIDFile(drive string) (bool, string) {
	uuidPath := filepath.Join(drive, "info", "uuid.txt")
	if checkFileExist(uuidPath) {
		content, err := readTextFile(uuidPath)
		if err != nil {
			Logger.Debug("Error:", err)
		} else {
			found, uuid := findUUID(content)
			if found {
				Logger.Debug("Found UUID:", uuid)
				return true, uuid
			}
		}
	}
	return false, ""
}

func findINFOFile(drive string) (bool, string) {
	infoPath := filepath.Join(drive, "info", "info.txt")
	if checkFileExist(infoPath) {
		content, err := readTextFile(infoPath)
		if err != nil {
			Logger.Debug("Error:", err)
		} else {
			if isHasName(content) {
				Logger.Debug("Found info")
				return true, content
			}
		}
	}
	return false, ""
}

func FindMatchedArcherDevice(drive string) bool {
	is_uuid, uuid := findUUIDFile(drive)
	is_info, info := findINFOFile(drive)
	if is_uuid && is_info {
		Logger.Info("FOUND Archer device: ", uuid)
		Logger.Info(info)
		return true
	}
	Logger.Debugf("Device on %s not match Archer device", drive)
	return false
}

func Walk(root string) []string {
	var paths []string

	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			Logger.Debug("Directory: ", path)
		} else {
			Logger.Debug("File: ", path)
			paths = append(paths, path)
		}
		return nil
	})
	if err != nil {
		Logger.Debug("Error:", err)
	}
	return paths
}

// GetFileExtension returns the file extension for a given file path.
func GetFileExtension(path string) string {
	return strings.ToLower(filepath.Ext(path))
}

// IsExtensionAllowed checks if the file extension is in the allowed list.
func IsExtensionAllowed(path string) bool {
	ext := GetFileExtension(path)
	for _, allowedExt := range allowedExtensions {
		if ext == allowedExt {
			return true
		}
	}
	return false
}

// RenameFile renames a file from oldPath to newPath.
func RenameFile(oldPath, newPath string) error {
	// Use os.Rename to rename the file.
	err := os.Rename(oldPath, newPath)
	if err != nil {
		return fmt.Errorf("failed to rename file: %w", err)
	}
	return nil
}

// Rename not allowed files in path
func RenameNotAllowed(dirPath string) {

	is_match := FindMatchedArcherDevice(dirPath)
	if is_match {
		files := Walk(dirPath)

		ret := false

		for _, f := range files {
			if !IsExtensionAllowed(f) {
				ret = PopupDoRename(dirPath, files)
			}
		}
		if ret {
			for _, f := range files {
				if !IsExtensionAllowed(f) {
					if RenameFile(f, f+".bak") == nil {
						Logger.Infof("Renamed: %s to %s", f, f+".bak")
					}
				}
			}
		}
	}

}

// Rename not allowed files in path
func ArchiveNotAllowed(dirPath string) error {

	is_match := FindMatchedArcherDevice(dirPath)
	if is_match {
		files := Walk(dirPath)

		for _, f := range files {
			if !IsExtensionAllowed(f) {
				if PopupDoArchive(dirPath, files) {
					return ArchiveFiles(
						dirPath,
						path.Join(dirPath, "unexpected.zip"),
						IsExtensionAllowed,
						true,
					)
				}
				break
			}
		}
	}
	return nil
}
