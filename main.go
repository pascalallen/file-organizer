package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func OrganizeFiles(directory string) {
	files, err := os.ReadDir(directory)
	if err != nil {
		log.Fatalf("Error reading directory: %v", err)
	}

	for _, file := range files {
		if file.IsDir() {
			continue
		}

		ext := strings.ToLower(filepath.Ext(file.Name()))
		if ext == "" {
			continue
		}

		extDir := filepath.Join(directory, ext[1:])
		if _, err := os.Stat(extDir); os.IsNotExist(err) {
			os.Mkdir(extDir, 0755)
		}

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
	if len(os.Args) < 2 {
		log.Fatalf("Please provide a directory to organize. Usage: %s <directory>", os.Args[0])
	}

	directory := os.Args[1]

	if _, err := os.Stat(directory); os.IsNotExist(err) {
		log.Fatalf("Directory does not exist: %s", directory)
	}

	OrganizeFiles(directory)
}
