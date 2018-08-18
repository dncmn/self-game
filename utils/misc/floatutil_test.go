package misc

import "testing"

func TestFloatEqual(t *testing.T) {
	var a float64 = 5.33333333333
	var b float64 = 5.33333322222
	if FloatEquals(a, b) {
		t.Error("Cannot be equal")
	}

	var c float64 = 5.33333333333
	var d float64 = 5.33333333333
	if !FloatEquals(c, d) {
		t.Error("Can be equal")
	}

	var e float64 = 1.61111
	var f float64 = 1.722
	var g float64 = 1.61
	if FloatCompare(e, f) != -1 {
		t.Error("Should be -1")
	}
	if FloatCompare(f, e) != 1 {
		t.Error("Should be 1")
	}
	if FloatCompare(g, e) == 0 {
		t.Error("Should be 0")
	}
}
