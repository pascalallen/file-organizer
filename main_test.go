package main

import (
	"os"
	"path/filepath"
	"testing"
)

func TestOrganizeFiles(t *testing.T) {
	dir, err := os.MkdirTemp("", "temp")
	if err != nil {
		t.Fatalf("Unable to create temporary directory: %v", err)
	}
	defer os.RemoveAll(dir)

	txtFile := "test.txt"
	csvFile := "sample.csv"
	otherFile := "other"

	readFiles := []string{txtFile, csvFile, otherFile}

	for _, file := range readFiles {
		if err := os.WriteFile(filepath.Join(dir, file), []byte(file), 0644); err != nil {
			t.Fatalf("Unable to write temporary files: %v", err)
		}
	}

	OrganizeFiles(dir)

	fileExist := func(dir, filename string) bool {
		info, err := os.Stat(filepath.Join(dir, filename))
		if err != nil {
			return false
		}
		return !info.IsDir()
	}

	t.Run("OrganizeTxtFiles", func(t *testing.T) {
		if !fileExist(filepath.Join(dir, "txt"), txtFile) {
			t.Errorf("Text file not moved to the correct directory")
		}
	})

	t.Run("OrganizeCsvFiles", func(t *testing.T) {
		if !fileExist(filepath.Join(dir, "csv"), csvFile) {
			t.Errorf("CSV file not moved to the correct directory")
		}
	})

	t.Run("LeaveNonExtensionFiles", func(t *testing.T) {
		if !fileExist(dir, otherFile) {
			t.Errorf("File without extension should not be moved")
		}
	})
}
