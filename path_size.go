package code

import (
	"fmt"
	"math"
	"os"
	"path/filepath"
	"strings"
)

func GetSize(path string, findHidden bool, recursive bool) (int, error) {

	file, err := os.Lstat(path)
	if err != nil {
		return 0, err
	}

	isDir := file.IsDir()
	if !isDir {
		return int(file.Size()), nil
	}

	files, err := os.ReadDir(path)
	if err != nil {
		return 0, err
	}

	sizeSum := 0
	for _, v := range files {
		fileInfo, err := v.Info()
		if err != nil {
			return 0, err
		}

		isDir := fileInfo.IsDir()

		if isDir && recursive {
			result, err := GetSize(filepath.Join(path, fileInfo.Name()), findHidden, recursive)

			if err != nil {
				return 0, err
			}
			sizeSum += result
			continue
		} else if isDir {
			continue
		}

		fileName := fileInfo.Name()
		isHidden := strings.HasPrefix(fileName, ".")

		if isHidden && !findHidden {
			continue
		}

		sizeSum += int(fileInfo.Size())
	}
	return sizeSum, nil
}

func FormatSize(size int, human bool) string {
	if !human {
		return fmt.Sprintf("%dB", size)
	}

	var floatSize = float64(size)
	var unit string
	var formatSize float64

	if floatSize >= math.Pow(float64(1024), float64(6)) {
		formatSize = floatSize / math.Pow(float64(1024), float64(6))
		unit = "EB"
	} else if floatSize >= math.Pow(float64(1024), float64(5)) {
		formatSize = floatSize / math.Pow(float64(1024), float64(5))
		unit = "PB"
	} else if floatSize >= math.Pow(float64(1024), float64(4)) {
		formatSize = floatSize / math.Pow(float64(1024), float64(4))
		unit = "TB"
	} else if floatSize >= math.Pow(float64(1024), float64(3)) {
		formatSize = floatSize / math.Pow(float64(1024), float64(3))
		unit = "GB"
	} else if floatSize >= math.Pow(float64(1024), float64(2)) {
		formatSize = floatSize / math.Pow(float64(1024), float64(2))
		unit = "MB"
	} else if floatSize >= 1024 {
		formatSize = floatSize / float64(1024)
		unit = "KB"
	} else {
		formatSize = floatSize
		unit = "B"
	}

	return fmt.Sprintf("%.1f%s", formatSize, unit)
}

func GetPathSize(path string, human, findHidden, recursive bool) (string, error) {
	sizeInt, err := GetSize(path, findHidden, recursive)
	if err != nil {
		return "0B", err
	}

	return FormatSize(sizeInt, human), nil
}
