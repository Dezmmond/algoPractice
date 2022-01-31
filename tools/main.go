package tools

import "fmt"
import "time"
import "math/rand"

func SliceGenerator(N int) []int {
	rand.Seed(time.Now().Unix())
	return rand.Perm(N)

	// TODO: test for errors with a "true random" slice
	// coming soon :)
	//slice := make([]int, size, size)
	//rand.Seed(time.Now().UnixNano())
	//for i := 0; i < size; i++ {
	//	slice[i] = rand.Intn(999) - rand.Intn(999)
	//}
	//return slice

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

func Pow(num int, power int) int {
	var result = 1
	for i := 0; i < power; i++ {
		result *= num
	}
	return result
}

func Checker(slice []int) bool {
	for i := 0; i < len(slice)-2; i++ {
		if slice[i] > slice[i+1] {
			fmt.Printf("This array was sorted incorrectly!\n"+
				"Indexes of the elements that led to the error: [%d : %d]\n", i, i+1)
			return false
		}
	}
	return true
}
