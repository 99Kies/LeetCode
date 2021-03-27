package main

import "fmt"

func busyStudent(startTime []int, endTime []int, queryTime int) int {
	length := len(startTime)
	sum := 0
	for i := 0; i < length; i++ {
		if startTime[i] <= queryTime && endTime[i] >= queryTime {
			sum++
		}
	}
	return sum
}

func main() {
	//startTime := []int{1,2,3}
	//endTime := []int{3,2,7}
	startTime := []int{4}
	endTime := []int{4}
	queryTime := 4
	a := busyStudent(startTime, endTime, queryTime)
	fmt.Println(a)
}
