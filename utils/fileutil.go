package utils

import (
	"os"
)

func FileExist(path string) bool {
	f, err := os.Open(path)
	if err != nil {
		return false
	}
	defer f.Close()
	return true
}
