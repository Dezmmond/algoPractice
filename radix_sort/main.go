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

func flatten(slice [][]int) (result []int) {
	for i := 0; i < len(slice); i++ {
		elem := slice[i]
		result = append(result, elem...)
	}
	return result
}

func main() {
	S := tools.SliceGenerator(10)
	digits := numDigits(S)
	for i := 0; i < digits; i++ {
		buff := make([][]int, 10)
		for _, elem := range S {
			num := (elem / tools.Pow(10, i)) % 10
			buff[num] = append(buff[num], elem)
		}
		S = flatten(buff)
	}
	fmt.Println(tools.Checker(S))
}
