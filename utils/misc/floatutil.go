package misc

import (
	"math/big"
)

const (
	EPSILON float64 = 0.00000001
)

func FloatEquals(a, b float64) bool {
	if (a-b) < EPSILON && (b-a) < EPSILON {
		return true
	}
	return false
}

func FloatCompare(a, b float64) int {
	x, y := big.NewFloat(a), big.NewFloat(b)
	return x.Cmp(y)
}
