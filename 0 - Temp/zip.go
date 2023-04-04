package main

import (
	"archive/zip"
	"io"
	"os"
	"path/filepath"
	"time"
)

func main() {
	sourceDir := "/home/tiagoreale/Imagens"
	destDir := "/home/tiagoreale/Downloads/GO_zip"
	interval := 24 * time.Hour // intervalo de 24 horas

	for {
		err := backup(sourceDir, destDir)
		if err != nil {
			// tratamento de erro
		}
		time.Sleep(interval)
	}
}

func backup(sourceDir, destDir string) error {
	backupFile, err := os.Create(filepath.Join(destDir, time.Now().Format("2006-01-02")+".zip"))
	if err != nil {
		return err
	}
	defer backupFile.Close()

	zipWriter := zip.NewWriter(backupFile)
	defer zipWriter.Close()

	err = filepath.Walk(sourceDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() {
			return nil
		}

		sourceFile, err := os.Open(path)
		if err != nil {
			return err
		}
		defer sourceFile.Close()

		destPath := path[len(sourceDir)+1:]
		destFile, err := zipWriter.Create(destPath)
		if err != nil {
			return err
		}

		_, err = io.Copy(destFile, sourceFile)
		if err != nil {
			return err
		}

		return nil
	})

	return err
}
