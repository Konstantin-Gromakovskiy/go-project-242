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
		flag     bool
		pathSize int
		err      error
	}{
		{"empty file", "testData/emptyFile.txt", false, 0, nil},
		{"not empty file", "testData/text.txt", false, 1094, nil},
		{"empty directory", "testData/emptyDir", false, 0, nil},
		{"not empty directory", "testData/notEmptyDir", false, 14, nil},
		{"not empty directory, search all", "testData/notEmptyDir", true, 55, nil},
	}

	for _, tt := range tests {
		fixturePath := filepath.Join(workDir, tt.path)
		t.Run(tt.name, func(t *testing.T) {
			funcResult, err := code.GetSize(fixturePath, tt.flag)
			assert.Equal(t, tt.err, err)
			assert.Equal(t, tt.pathSize, funcResult)
		})
	}
}
