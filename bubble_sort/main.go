package main

import "algoPractice/tools"

func main() {
	shuffled := tools.SliceGenerator(10)
	sliceCopy := make([]int, len(shuffled))
	copy(sliceCopy, shuffled)

	for j := 1; j < len(shuffled); j++ {
		go_on := false
		for i := 0; i < len(shuffled)-j; i++ {
			if shuffled[i] > shuffled[i+1] {
				shuffled[i], shuffled[i+1] = shuffled[i+1], shuffled[i]
				go_on = true
			}
		}
		if !go_on {
			break
		}
	}
	tools.Checker("Bubble Sort", true, sliceCopy, shuffled)
}
