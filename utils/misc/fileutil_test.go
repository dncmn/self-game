package misc

import (
	"testing"
)

func TestFileExist(t *testing.T) {
	// check self is exist
	if !FileExist("./fileutil_test.go") {
		t.Error("cannot be not exist")
	}

	// check self folder is exist
	if !FileExist("/home") {
		t.Error("cannot be not exist")
	}
}
