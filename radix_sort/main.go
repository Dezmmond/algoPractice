package main

import "fmt"
import "algoPractice/tools"

// Get number of digits in largest element
func numDigits(slice []int) int {
	var maxElem int
	var digits = 1
	var del = 10
	for _, elem := range slice {
		maxElem = tools.Max(maxElem, elem)
	}
	for maxElem/del > 0 {
		del *= 10
		digits += 1
	}
	return digits
}

func main() {
	S := tools.SliceGenerator(1001)
	max := numDigits(S)
	fmt.Println(max)
}
