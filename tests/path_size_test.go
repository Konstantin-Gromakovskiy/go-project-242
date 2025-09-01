package tests

import (
	"code"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetPathSize_File(t *testing.T) {
	workDir, _ := os.Getwd()

	tests := []struct {
		name          string
		path          string
		hiddenFlag    bool
		recursiveFlag bool
		pathSize      int
		err           error
	}{
		{"empty file", "testData/emptyFile.txt", false, false, 0, nil},
		{"not empty file", "testData/text.txt", false, false, 1094, nil},
		{"empty directory", "testData/emptyDir", false, false, 0, nil},
		{"not empty directory", "testData/notEmptyDir", false, false, 14, nil},
		{"not empty directory, search all", "testData/notEmptyDir", true, false, 55, nil},
		{"recursive directory", "testData/recursiveDir", false, false, 14, nil},
		{"recursive dir, search all", "testData/recursiveDir", true, false, 55, nil},
		{"recursive dir, search all, recursive", "testData/recursiveDir", true, true, 110, nil},
		{"recursive dir, search visible, recursive", "testData/recursiveDir", false, true, 28, nil},
		{"wrong path", "wrong/path", false, false, 0, assert.AnError},
	}

	for _, tt := range tests {
		fixturePath := filepath.Join(workDir, tt.path)
		t.Run(tt.name, func(t *testing.T) {
			funcResult, err := code.GetPathSize(fixturePath, tt.hiddenFlag, tt.recursiveFlag)
			if tt.err != nil {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
			assert.Equal(t, tt.pathSize, funcResult)
		})
	}
}
