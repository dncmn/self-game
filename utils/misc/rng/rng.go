package rng

import (
	"math/rand"
	"time"
)

func Seed(seed int64) {
	rand.Seed(seed)
}

func GetSeed() int64 {
	seed := int64(0)
	s := rand.NewSource(0)
	N := 10
	for i := 0; i < N; i++ {
		t := time.Now().UnixNano()
		s.Seed(t)
		seed ^= s.Int63()
	}
	return seed
}

func Next() int32 {
	return rand.Int31n(10000)
}

func IntN(v uint64) int32 {
	return int32(rand.Int63n(int64(v)))
}

func IntRange(min, max uint64) uint64 {
	if min >= max {
		return min
	}
	v := rand.Int63n(int64(max - min))
	return min + uint64(v)
}

func Perm(n int) []int {
	return rand.Perm(n)
}

func NormalVariate(mean float64, stdDev float64) float64 {
	return rand.NormFloat64()*stdDev + mean
}
