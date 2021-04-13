package main

import "fmt"

func minOperations(boxes string) []int {
	//result := make([]int, len(boxes))
	length := len(boxes)
	result := make([]int, length)
	//fmt.Println(length)
	//var return_result []int
	pre := 0
	sum := 0
	for i := 0; i < length; i++ {
		result[i] += sum
		if boxes[i] == '1' {
			pre++
		}
		sum += pre
	}

	pre = 0
	sum = 0
	for i := length - 1; i >= 0; i-- {
		result[i] += sum
		if boxes[i] == '1' {
			pre++
		}
		sum += pre
	}
	return result
}

func main() {
	boxes := "110"
	res := minOperations(boxes)
	fmt.Println(res)
}
