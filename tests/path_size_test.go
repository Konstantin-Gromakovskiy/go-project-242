package tests

import (
	code "code/src"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetPathSize_File(t *testing.T) {
	workDir, _ := os.Getwd()

	tests := []struct {
		name     string
		path     string
		pathSize int
		err      error
	}{
		{"empty file", "testData/emptyFile.txt", 0, nil},
		{"not empty file", "testData/text.txt", 13, nil},
		{"empty directory", "testData/emptyDir", 0, nil},
		{"not empty directory", "testData/notEmptyDir", 14, nil},
	}

	for _, tt := range tests {
		fixturePath := filepath.Join(workDir, tt.path)
		t.Run(tt.name, func(t *testing.T) {
			funcResult, err := code.GetSize(fixturePath)
			assert.Equal(t, tt.err, err)
			assert.Equal(t, tt.pathSize, funcResult)
		})
	}
}
