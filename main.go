package main

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func organizeFiles(directory string) {
	files, err := os.ReadDir(directory)
	if err != nil {
		log.Fatalf("Error reading directory: %v", err)
	}

	for _, file := range files {
		if file.IsDir() || strings.HasPrefix(file.Name(), ".") {
			continue
		}

		ext := strings.ToLower(filepath.Ext(file.Name()))
		if ext == "" {
			continue
		}

		extDir := filepath.Join(directory, ext[1:])
		if _, err := os.Stat(extDir); os.IsNotExist(err) {
			err := os.Mkdir(extDir, 0755)
			if err != nil {
				log.Printf("Failed to create directory %s: %v", extDir, err)
				continue
			}
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
	myApp := app.New()
	myWindow := myApp.NewWindow("File Organizer")
	myWindow.Resize(fyne.NewSize(500, 300))

	entry := widget.NewEntry()
	entry.SetPlaceHolder("Enter directory path")

	output := widget.NewLabel("")

	btn := widget.NewButton("Organize", func() {
		dir := entry.Text
		if _, err := os.Stat(dir); os.IsNotExist(err) {
			output.SetText(fmt.Sprintf("Directory does not exist: %s", dir))
			return
		}
		organizeFiles(dir)
		output.SetText("Files organized successfully")
	})

	myWindow.SetContent(
		container.NewVBox(
			widget.NewLabel("File Organizer"),
			entry,
			btn,
			output,
		),
	)

	myWindow.ShowAndRun()
}
