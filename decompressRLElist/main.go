package main

import "fmt"

func decompressRLElist(nums []int) []int {
	length := len(nums)
	var return_result []int
	var flag_result []int

	for i := 0; i < length/2; i++ {
		flag_result = make([]int, nums[2*i])
		for j := 0; j < nums[2*i]; j++ {
			flag_result[j] = nums[2*i+1]

		}
		return_result = append(return_result, flag_result...)
		flag_result = flag_result[:]
	}

	return return_result

}

func main() {
	a := []int{1, 2, 3, 4}
	b := decompressRLElist(a)
	fmt.Println()
	fmt.Println(b)
}
