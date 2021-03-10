package main

import "fmt"

func minCount(coins []int) int {
	result := 0
	for i := range coins {
		result = result + coins[i]/2 + coins[i]%2
	}
	return result
}

func main() {
	a := []int{4, 2, 1}
	b := []int{2, 3, 10}
	fmt.Println(minCount(a))
	fmt.Println(minCount(b))
}
