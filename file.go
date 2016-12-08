package vhelper

import "os"

func IsExist(fn string) bool {
	f, err := os.Open(fn)
	defer f.Close()
	if err != nil && os.IsNotExist(err) {
		return false
	}
	return true
}
