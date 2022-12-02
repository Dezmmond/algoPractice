package main

import "algoPractice/tools"

// TODO: Work correctly only with positive numbers

// Get number of digits in largest element
func numDigits(slice []int) int {
	var maxElem int
	var numLength, divider = 1, 10

	maxElem = tools.GetMaximum(slice)

	for maxElem/divider > 0 {
		divider *= 10
		numLength++
	}
	return numLength
}

// Forming a new slice from the buffer
func flatten(buffer [][]int) (result []int) {
	for i := 0; i < len(buffer); i++ {
		elem := buffer[i]
		result = append(result, elem...)
	}
	return result
}

func main() {
	shuffled := tools.SliceGenerator(10)
	sliceCopy := make([]int, len(shuffled))
	copy(sliceCopy, shuffled)

	digits := numDigits(shuffled)
	for i := 0; i < digits; i++ {
		// Creating a slice of slices for storing numbers by digits
		buffer := make([][]int, 10)
		for _, num := range shuffled {
			digit := (num / tools.Pow(10, i)) % 10
			if digit < 0 {
				digit *= -1
			}
			buffer[digit] = append(buffer[digit], num)
		}
		shuffled = flatten(buffer)
	}
	tools.Checker("Radix Sort", true, sliceCopy, shuffled)
}
