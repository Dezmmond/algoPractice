package main

import "algoPractice/tools"

func main() {
	shuffled := tools.SliceGenerator(12)
	sliceCopy := make([]int, len(shuffled))
	copy(sliceCopy, shuffled)

	for i := 1; i < len(shuffled); i++ {
		x := i
		for shuffled[x] < shuffled[x-1] {
			shuffled[x-1], shuffled[x] = shuffled[x], shuffled[x-1]
			if x-1 == 0 {
				break
			}
			x--
		}

	}
	tools.Checker("Insertion Sort", true, sliceCopy, shuffled)
}
