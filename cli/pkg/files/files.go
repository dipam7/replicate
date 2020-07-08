package files

import (
	"os"
)

func FileExists(filePath string) (bool, error) {
	if _, err := os.Stat(filePath); err == nil {
		return true, nil
	} else if os.IsNotExist(err) {
		return false, nil
	} else {
		return false, err
	}
}

func IsDir(dirPath string) (bool, error) {
	file, err := os.Stat(dirPath)
	if err != nil {
		return false, err
	}
	return file.Mode().IsDir(), nil
}