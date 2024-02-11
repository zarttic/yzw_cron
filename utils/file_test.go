package utils

import (
	"os"
	"path/filepath"
	"testing"
)

func TestFile(t *testing.T) {
	body := []byte("test body")
	err := File(body)
	if err != nil {
		t.Errorf("Error writing file: %v", err)
	}

	// Verify that the file was created
	filePath := TimeStr("2006-01-02-15-04")
	dirPath := "craw/" + TimeStr("2006-01-02")
	filePath = filepath.Join(dirPath, filePath+".txt")
	_, err = os.Stat(filePath)
	if os.IsNotExist(err) {
		t.Errorf("File not created: %s", filePath)
	}
}
