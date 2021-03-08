package main

import "fmt"

func game(guess []int, answer []int) int {
	result := 0
	for i := range guess {
		if guess[i] == answer[i] {
			result++
		}
	}
	return result
}

func main() {
	guess_1 := []int{1, 2, 3}
	answer_1 := []int{1, 1, 3}
	fmt.Println(game(guess_1, answer_1))
}
