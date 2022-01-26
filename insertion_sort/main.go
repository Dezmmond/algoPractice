package main

import "fmt"

func main() {
	S := []int{6, 5, 3, 12, 8, 7, 2, 4, 10, 11, 9, 1}

	for i := 1; i < len(S); i++ {
		x := i
		for S[x] < S[x-1] {
			S[x-1], S[x] = S[x], S[x-1]
			if x-1 == 0 {
				break
			}
			x--
		}

	}
	fmt.Print(S)
}
