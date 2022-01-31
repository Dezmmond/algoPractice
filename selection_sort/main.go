package main

import "fmt"
import "algoPractice/tools"

func main() {
	S := tools.SliceGenerator(10)
	for i := 0; i < len(S); i++ {
		min := i
		for j := i + 1; j < len(S); j++ {
			if S[j] < S[min] {
				min = j
			}
		}
		S[i], S[min] = S[min], S[i]
	}
	fmt.Println(tools.Checker(S))
}
