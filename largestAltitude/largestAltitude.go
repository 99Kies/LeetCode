package main

import "fmt"

func largestAltitude(gain []int) int {
	length := len(gain)
	start := 0
	max := 0
	for i := 0; i < length; i++ {
		next := gain[i] + start
		if next > max {
			max = next
		}
		start = next
	}
	return max
}

func main() {
	p := []int{-4, -3, -2, -1, 4, 3, 2}
	a := largestAltitude(p)
	fmt.Println(a)
}
