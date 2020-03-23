package fileutils

import (
	"io/ioutil"
	"os"
)

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

func IsDir(pathname string) bool {
	s, err := os.Stat(pathname)
	if err != nil {
		return false
	}
	return s.IsDir()
}

func MkDir(dirname string) error {

	return os.MkdirAll(dirname, os.ModePerm)

}

func MkDirIfNotExists(dirname string) error {

	if IsDir((dirname)) {
		return nil
	}

	return MkDir(dirname)

}

func Remove(pathname string) error {

	return os.Remove(pathname)
}

func ReadFile(pathname string) []byte {

	bs, _ := ioutil.ReadFile(pathname)

	return bs
}
