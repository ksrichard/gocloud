package util

import "os"

func IsDir(dirInput string) bool {
	fi, err := os.Stat(dirInput)
	if err != nil {
		return false
	}
	switch mode := fi.Mode(); {
	case mode.IsDir():
		return true
	case mode.IsRegular():
		return false
	}

	return false
}

func CreateDir(folderPath string) error {
	if IsDir(folderPath) {
		return nil
	}
	return os.MkdirAll(folderPath, os.ModePerm)
}
