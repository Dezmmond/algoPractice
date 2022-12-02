package main

import "algoPractice/tools"

func main() {
	shuffled := tools.SliceGenerator(40)
	result := quickSort(shuffled)
	tools.Checker("Quick Sort", true, shuffled, result)
}

func quickSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}

	var left []int
	var right []int

	a, b, c := arr[0], arr[len(arr)/2], arr[len(arr)-1]
	pivot := (a + b + c) - (tools.Max(tools.Max(a, b), c) + tools.Min(tools.Min(a, b), c))

	for _, elem := range arr {
		if elem < pivot {
			left = append(left, elem)
		} else if elem > pivot {
			right = append(right, elem)
		}
	}
	return append(append(quickSort(left), pivot), quickSort(right)...)
}
