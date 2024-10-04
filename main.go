package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

// OrganizeFiles moves files into directories based on their extensions
func OrganizeFiles(directory string) {
	// Read the contents of the directory
	files, err := ioutil.ReadDir(directory)
	if err != nil {
		log.Fatalf("Error reading directory: %v", err)
	}

	// Iterate over each file in the directory
	for _, file := range files {
		// Skip if it's a directory
		if file.IsDir() {
			continue
		}

		// Get the file extension (lowercased)
		ext := strings.ToLower(filepath.Ext(file.Name()))
		if ext == "" {
			continue // Skip files without an extension
		}

		// Create a directory for the extension if it doesn't exist
		extDir := filepath.Join(directory, ext[1:]) // Removing the leading dot from extension
		if _, err := os.Stat(extDir); os.IsNotExist(err) {
			os.Mkdir(extDir, 0755)
		}

		// Move the file into the appropriate directory
		oldPath := filepath.Join(directory, file.Name())
		newPath := filepath.Join(extDir, file.Name())
		if err := os.Rename(oldPath, newPath); err != nil {
			log.Printf("Failed to move file %s: %v", file.Name(), err)
		} else {
			fmt.Printf("Moved %s to %s\n", file.Name(), newPath)
		}
	}
}

func main() {
	// Check if the user provided a directory path
	if len(os.Args) < 2 {
		log.Fatalf("Please provide a directory to organize. Usage: %s <directory>", os.Args[0])
	}

	// Get the directory from the command line arguments
	directory := os.Args[1]

	// Check if the directory exists
	if _, err := os.Stat(directory); os.IsNotExist(err) {
		log.Fatalf("Directory does not exist: %s", directory)
	}

	// Organize files in the specified directory
	OrganizeFiles(directory)
}
