package utils

import (
	"archive/zip"
	"io"
	"os"
	"path/filepath"
)

func ArchiveFiles(sourceDir string, destFilePath string, filter func(string) bool, removeArchived bool) error {
	// Create a new zip archive file
	newArchive, err := os.Create(destFilePath)
	if err != nil {
		return err
	}
	defer newArchive.Close()

	// Create a new zip writer
	zipWriter := zip.NewWriter(newArchive)
	defer zipWriter.Close()

	// Walk through the source directory recursively
	err = filepath.Walk(sourceDir, func(filePath string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// Skip directories
		if info.IsDir() {
			return nil
		}

		// Append filter
		if filter != nil {
			if filter(filePath) {
				return nil
			}
		}

		// Create a new file in the zip archive corresponding to the file's relative path
		relPath, err := filepath.Rel(sourceDir, filePath)
		if err != nil {
			return err
		}

		// Open the source file
		sourceFile, err := os.Open(filePath)
		if err != nil {
			return err
		}
		defer sourceFile.Close()

		// Create a new file in the zip archive
		destFile, err := zipWriter.Create(relPath)
		if err != nil {
			return err
		}

		// Copy the contents of the source file to the destination file in the zip archive
		_, err = io.Copy(destFile, sourceFile)
		if err != nil {
			return err
		}

		Logger.Debugf("Archived: %s", filePath)

		if removeArchived {
			// Close the file before attempting to remove it
			if err := sourceFile.Close(); err != nil {
				return err
			}
			Logger.Debugf("Closed: %s", filePath)

			// Attempt to remove the file
			if err := os.Remove(filePath); err != nil {
				return err
			}
			Logger.Debugf("Removed: %s", filePath)
		}

		return nil
	})

	return err
}

// func main() {
// 	sourceDir := "../tmp"
// 	destFilePath := "unexpected.zip"
// 	err := archiveFiles(sourceDir, destFilePath)
// 	if err != nil {
// 		panic(err)
// 	}
// 	println("Files archived successfully!")
// }
