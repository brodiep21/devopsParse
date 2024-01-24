package organize

import (
	"io"
	"os"
	"path/filepath"
)

func OrganizeFiles(sourceDir, destDir string) error {
	err := filepath.Walk(sourceDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil
		}

		fileType := getFileType(info.Name())

		destFolder := filepath.Join(destDir, fileType)
		if err := os.MkdirAll(destFolder, os.ModePerm); err != nil {
			return err
		}

		destPath := filepath.Join(destFolder, info.Name())
		return moveFile(path, destPath)
	})

	return err
}

func getFileType(fileName string) string {
	ext := filepath.Ext(fileName)
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

func moveFile(source, destination string) error {
	sourceFile, err := os.Open(source)
	if err != nil {
		return err
	}

	defer sourceFile.Close()

	destFile, err := os.Create(destination)
	if err != nil {
		return err
	}
	defer destFile.Close()

	_, err = io.Copy(destFile, sourceFile)
	if err != nil {
		return err
	}

	return os.Remove(source)
}
