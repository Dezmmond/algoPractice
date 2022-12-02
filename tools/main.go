package tools

import (
	"fmt"
	"math/rand"
	"time"
)

func SliceGenerator(size int) []int {
	if size < 0 {
		size *= -1
	}
	slice := make([]int, size, size)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < size; i++ {
		slice[i] = rand.Intn(999) - rand.Intn(599)
	}
	return slice
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

func GetMaximum(slice []int) int {
	var maximum = 0
	for _, elem := range slice {
		maximum = Max(maximum, elem)
	}
	return maximum
}

func Pow(num int, power int) int {
	var result = 1
	for i := 0; i < power; i++ {
		result *= num
	}
	return result
}

func Checker(sortName string, verbose bool, original, result []int) bool {
	if verbose {
		fmt.Printf("Original slice: %v\n"+
			"Result of '%s' algorithm: %v\n", original, sortName, result)
	}
	if len(original) != len(result) {
		fmt.Println("The number of elements in the slices is not equal")
		return false
	}
	for i := 0; i < len(result)-2; i++ {
		if result[i] > result[i+1] {
			fmt.Printf("This array was sorted incorrectly!\n"+
				"Indexes of the elements that led to the error: [%d : %d]\n", i, i+1)
			return false
		}
	}
	return true
}
