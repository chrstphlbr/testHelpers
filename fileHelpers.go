package testHelpers

import (
	"os"
	"testing"
)

func CreateFile(t *testing.T, fileName, content string) {
	file, err := os.Create(fileName)
	if err != nil {
		t.Fatalf("could not create file: %v", err)
	}
	_, err = file.WriteString(content)
	if err != nil {
		t.Fatalf("could not write file: %v", err)
	}
	file.Close()
}

func RemoveFile(t *testing.T, fileName string) {
	err := os.RemoveAll(fileName)
	if err != nil {
		t.Logf("could not delete file (path:\"%s\")", fileName)
	}
}
