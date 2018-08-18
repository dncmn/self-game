package conv

import (
	"fmt"
	"testing"
)

func TestIsInGB2312(t *testing.T) {
	str := "aaa三ae恶sddd俄341啊吧"
	if !IsAllGB2312(str) {
		t.Error("Error code")
	}

}

func TestIsHasSpace(t *testing.T) {
	str := " aaa三a"
	if !IsHasSpace(str) {
		t.Error("Error code")
	}
}

func TestHanZi(t *testing.T) {
	var x rune = '啊'
	i := 1
	for ; i <= 20188-1000-100-14; i++ {
		if !IsAllGB2312(string(rune(int(x) + i))) {
			t.Error("Error code")
		}
	}
}

func TestHanZiLength(t *testing.T) {
	str := "你好hello"
	fmt.Println(NameLength(str))
}

func TestSortMapByStringKey(t *testing.T) {
	vs := map[string]string{"1": "aaa", "0": "bbb", "3": "ccc"}
	for i, v := range vs {
		fmt.Println(i, v)
	}

	result := SortMapByStringKey(vs)
	for i, v := range result {
		fmt.Println(i, v)
	}
}

func TestSortMapByIntKey(t *testing.T) {
	vs := map[int]string{1: "aaa", 0: "bbb", 3: "ccc"}
	for i, v := range vs {
		fmt.Println(i, v)
	}

	result := SortMapByIntKey(vs)
	for i, v := range result {
		fmt.Println(i, v)
	}
}
