package main

import "fmt"

//func lengthOfLongestSubstring(s string) int {
//
//}

func main() {
	costs := []int{2, 7, 12}
	days := []int{1, 4, 6, 7, 8, 20}
	dpCount := days[len(days)-1] + 1
	dpArray := make([]int, dpCount)

	for _, day := range days {
		dpArray[day]++
	}
	for i := 1; i < dpCount; i++ {
		if dpArray[i] > 0 {
			dpArray[i] = min(
				dpArray[i-1]+costs[0],
				dpArray[max(i-7, 0)]+costs[1],
				dpArray[max(i-30, 0)]+costs[2],
			)
		} else {
			dpArray[i] = dpArray[i-1]
		}
	}
	fmt.Println(dpArray)

}

func min(a, b, c int) int {
	//result := 0
	if a < b {
		b = a
	} else {
		a = b
	}
	if a < c {
		c = a
	}
	return c
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}
