package utils

import (
	"fmt"
	"os/exec"
	"path/filepath"
)

// GetDriveLetter extracts the drive letter from a given Windows path.
func GetDriveLetter(path string) (string, error) {
	absPath, err := filepath.Abs(path)
	if err != nil {
		return "", err
	}

	if len(absPath) < 2 {
		return "", fmt.Errorf("invalid path: %s", path)
	}

	// Drive letter is the first two characters (e.g., "C:").
	if absPath[1] == ':' {
		return absPath[:2], nil
	}

	return "", fmt.Errorf("no drive letter found in path: %s", path)
}

func EjectDrive(path string) error {

	if letter, err := GetDriveLetter(path); err != nil {
		fmt.Println(err)
		return err
	} else {

		command := `(New-Object -comObject Shell.Application).Namespace(17).ParseName("%s").InvokeVerb("Eject")`

		command = fmt.Sprintf(command, letter)

		// Create the PowerShell command
		cmd := exec.Command("powershell", "-Command", command)

		// Run the command and capture the output
		_, err := cmd.CombinedOutput()
		if err != nil {
			Logger.Errorf("Failed to execute command: %s", err)
			return err
		}

		Logger.Infof("Ejected %s", letter)
		return nil
	}
}

// func main() {
// 	EjectDrive(`GolandProjects\archer-atichea`)
// }
