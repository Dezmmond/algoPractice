package tools

import "time"
import "math/rand"

func SliceGenerator(N int) []int {
	rand.Seed(time.Now().Unix())
	return rand.Perm(N)
}

func Min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}

func Max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
