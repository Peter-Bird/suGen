package gotree

import (
	"fmt"
	"io"
	"log"
	"os"
)

// CheckDirExists:
// - Verifies if a directory exists.
// Returns
// - nil if the directory does not exist,
// - an error if it does.
func CheckDirExists(directoryName string) error {
	_, err := os.Stat(directoryName)
	if err == nil {
		return fmt.Errorf("directory %q already exists", directoryName)
	}
	if !os.IsNotExist(err) {
		return err // Return the error if it's not a "does not exist" error
	}
	return nil // Return nil if the directory does not exist
}

// CreateDirs:
// Creates a series of directories based on the provided slice of paths.
// It returns an error if any directory creation fails.
func CreateDirs(dirs []string, perm os.FileMode, logger func(string)) error {
	for _, dir := range dirs {
		err := os.MkdirAll(dir, perm)
		if err != nil {
			return fmt.Errorf("error creating directory %s: %w", dir, err)
		}
		logger(fmt.Sprintf("- Directory %s created.\n", dir))
	}
	return nil
}

func CreateFiles(files map[string]string, logger func(string)) {
	for fileName, content := range files {
		if err := MakeFile(fileName, content, logger); err != nil {
			log.Fatalf("Error creating or writing to file %s: %s", fileName, err)
		}
	}
}

func MakeFile(fileName, content string, logger func(string)) error {
	file, err := os.Create(fileName)
	if err != nil {
		return fmt.Errorf("error creating file %s: %s", fileName, err)
	}
	defer file.Close()

	if _, err = io.WriteString(file, content); err != nil {
		return fmt.Errorf("error writing to file %s: %s", fileName, err)
	}

	logger(fmt.Sprintf("- File %s created and populated.\n", fileName))
	return nil
}
