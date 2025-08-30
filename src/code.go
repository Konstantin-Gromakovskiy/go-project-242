package code

import (
	"os"
)

func GetSize(path string) (int, error) {
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

		sizeSum += int(fileInfo.Size())
	}
	return sizeSum, nil
}
