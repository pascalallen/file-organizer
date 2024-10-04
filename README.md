# File Organizer

A simple Go application that organizes files into directories based on their file extensions. It moves files with the same extension into respective folders (e.g., `.txt` files will go into a `txt/` folder, `.jpg` files will go into a `jpg/` folder, and so on).

## Features

- Organizes files based on their extensions.
- Dynamically creates folders for each unique file extension.
- Allows users to specify a directory to organize via command-line arguments.

## Download(s)

You can download a precompiled version of the File Organizer for your operating system:

- [Download for Linux](https://github.com/pascalallen/file-organizer/releases/download/v1.0.0/file-organizer-linux)
- [Download for macOS](https://github.com/pascalallen/file-organizer/releases/download/v1.0.0/file-organizer-macos)
- [Download for Windows](https://github.com/pascalallen/file-organizer/releases/download/v1.0.0/file-organizer-windows.exe)

## Usage

Once downloaded, you can run the program from the command line, specifying the directory you want to organize.

### Linux/macOS:

```bash
./file-organizer-macos /path/to/your/directory
```

### Windows

```bash
file-organizer-windows.exe C:\path\to\your\directory
```