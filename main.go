package main

import (
	"fmt"
	"net/http"
)

func main() {
	sourceDir := "path/to/source"
	destDir := "path/to/dest"

	err := organizeFiles(sourceDir, destDir)
	if err != nil {
		fmt.Println("Error organizing files:", err)
	}


	// http.HandleFunc("/", handler)
	// http.ListenAndServe(":8081", nil)
	// fmt.Println("Hello World!")
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World")
}


func organizeFiles(sourcePath, destPath string) error {
	err := filePath.Walk(sourcePath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.isDir() {
			return nil
		}
	},

	// fileType := getFileType
	return err
}

func getFileType(fileName string) string {
	ext := filePath.Ext(fileName)
	switch ext {
	case ".jpg", ".png", ".webp":
		return "Images"
	case ".gif":
		return "Gifs"
	case ".pdf", ".doc", ".docx", ".txt":
		return "Documents"
	case ".mp4", ".avi", ".mkv":
		return "Videos"
	case ".mp3":
		return "Music"
	default:
		return "Other"
	}
}