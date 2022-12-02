package main

import "algoPractice/tools"

func main() {
	shuffled := tools.SliceGenerator(10)
	sliceCopy := make([]int, len(shuffled))
	copy(sliceCopy, shuffled)

	for i := 0; i < len(shuffled); i++ {
		min := i
		for j := i + 1; j < len(shuffled); j++ {
			if shuffled[j] < shuffled[min] {
				min = j
			}
		}
		shuffled[i], shuffled[min] = shuffled[min], shuffled[i]
	}
	tools.Checker("Selection Sort", true, sliceCopy, shuffled)
}
