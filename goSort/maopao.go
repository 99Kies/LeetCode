package main

import "fmt"

func maopao(nums []int) []int {
	for i := 0; i < len(nums)-1; i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i] > nums[j] {
				nums[i], nums[j] = nums[j], nums[i]
			}
		}
	}
	fmt.Println(nums)
	return nums
}

func maopaoSort(nums []int) []int {
	length := len(nums)
	for i := 0; i < length-1; i++ {
		for j := i + 1; j < length; j++ {
			if nums[i] > nums[j] {
				nums[i], nums[j] = nums[j], nums[i]
			}
		}
	}
	return nums
}

func main() {
	nums := []int{1, 2, 5, 3, 4}
	result := maopaoSort(nums)
	fmt.Println(result)
}
