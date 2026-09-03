package backend

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
)

func SaveToLocal(r io.Reader, objectName string) (string, error) {
	dir := "uploads"
	if err := os.MkdirAll(dir, 0755); err != nil {
		return "", err
	}

	path := filepath.Join(dir, objectName)
	f, err := os.Create(path)
	if err != nil {
		return "", err
	}
	defer f.Close()

	if _, err := io.Copy(f, r); err != nil {
		return "", err
	}

	url := "http://localhost:8080/uploads/" + objectName
	fmt.Printf("File is saved locally: %s\n", url)
	return url, nil
}