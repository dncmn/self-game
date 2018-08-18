package array

import (
	"testing"
)

func TestFindArray(t *testing.T) {
	array := []int32{1, 2, 3, 4, 5, 6, 7}
	if FindInt32InArray(array, 3) == -1 {
		t.Error("Error find array")
	}
	if FindInt32InArray(array, 8) != -1 {
		t.Error("Error find array")
	}
}

func TestAppendInt32ToArraySet(t *testing.T) {
	array := []int32{1, 2, 3, 4, 5, 6, 7}
	AppendInt32ToArraySet(&array, 8)
	t.Log(array)
}

func TestDeleteInt32FromArraySet(t *testing.T) {
	array := []int32{1, 2, 3, 4, 5, 6, 7}
	DelInt32FromArraySet(&array, 5)
	t.Log(array)
}
