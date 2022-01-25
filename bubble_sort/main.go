package main

import "fmt"

func main() {
	S := []int{10, 5, 14, 7, 3, 2, 18, 4, 5, 13, 6, 8}

	for j := 1; j < len(S); j++ {
		go_on := false
		for i := 0; i < len(S)-j; i++ {
			if S[i] > S[i+1] {
				S[i], S[i+1] = S[i+1], S[i]
				go_on = true
			}
		}
		if !go_on {
			break
		}
	}

	fmt.Println(S)
}
