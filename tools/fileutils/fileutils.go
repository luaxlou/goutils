package fileutils

import "os"

func Exists(pathname string) bool {

	_, err := os.Stat(pathname)
	if err == nil {

		return true

	}
	if os.IsExist(err) {
		return true
	}
	return false
}

func Remove(pathname string) error {

	return os.Remove(pathname)
}
