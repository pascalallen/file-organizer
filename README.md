# File Organizer

A simple Go application that organizes files into directories based on their file extensions. It moves files with the same extension into respective folders (e.g., `.txt` files will go into a `txt/` folder, `.jpg` files will go into a `jpg/` folder, and so on).

## Features

- Organizes files based on their extensions.
- Dynamically creates folders for each unique file extension.
- Allows users to specify a directory to organize via command-line arguments.

## Prerequisites

- [Go](https://golang.org/doc/install) (version 1.15+)
- A directory with files to organize.

## Installation

1. Clone the repository or download the source code:

    ```bash
    git clone https://github.com/pascalallen/file-organizer.git
    cd file-organizer
    ```

2. Initialize Go modules:

    ```bash
    go mod init file-organizer
    ```

## Usage

1. Run the program, specifying the directory you want to organize:

    ```bash
    go run main.go /path/to/your/directory
    ```

   For example, to organize files in the `downloads` folder in the current directory:

    ```bash
    go run main.go ./downloads
    ```

2. The application will:
    - Create folders based on file extensions (e.g., `txt/`, `jpg/`, `go/`).
    - Move files into their respective extension-based folders.

## Example

Before running the program, assume you have the following files in your `downloads` directory:

```bash
downloads/
│
├── document1.txt
├── image1.jpg
├── script.go
└── notes.txt
```

After running:

```bash
go run main.go ./downloads
```

The directory will be organized as:

```bash
downloads/
│
├── go/
│   └── script.go
├── jpg/
│   └── image1.jpg
└── txt/
    ├── document1.txt
    └── notes.txt
```

## Error Handling

If no directory is provided, the program will display an error message:

```bash
Please provide a directory to organize. Usage: ./file-organizer <directory>
```

## License

[MIT](LICENSE)
