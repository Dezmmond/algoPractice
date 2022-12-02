package main

import "algoPractice/tools"

func main() {
	shuffled := tools.SliceGenerator(12)
	result := mergeSort(shuffled)
	tools.Checker("Merge Sort", true, shuffled, result)
}

func merge(left []int, right []int) []int {
	var result []int
	i, j := 0, 0

	for i < len(left) && j < len(right) {
		if left[i] < right[j] {
			result = append(result, left[i])
			i++
		} else {
			result = append(result, right[j])
			j++
		}
	}
	if i < len(left) {
		result = append(result, left[i:]...)
	}
	if j < len(right) {
		result = append(result, right[j:]...)
	}
	return result
}

func mergeSort(S []int) []int {
	if len(S) == 1 {
		return S
	}
	mid := len(S) / 2
	left := mergeSort(S[:mid])
	right := mergeSort(S[mid:])

	return merge(left, right)
}
