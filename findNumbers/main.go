package main

import "fmt"

func findNumbers(nums []int) int {
	result := 0
	for i := range nums {
		flag := 1
		j := nums[i] / 10
		for j > 0 {
			//fmt.Println(j)
			j = j / 10
			flag++
		}
		if flag%2 == 0 {
			result++
		}
	}
	return result
}

func main() {
	num := []int{12, 345, 23, 7832}
	result := findNumbers(num)
	fmt.Println(result)
}
