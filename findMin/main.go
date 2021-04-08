package main

import "fmt"

//func findMin(nums []int) int {
//	low, high := 0, len(nums) - 1
//	for low < high {
//		pivot := low + (high - low) / 2
//		if nums[pivot] < nums[high] {
//			high = pivot
//		} else {
//			low = pivot + 1
//		}
//	}
//	return nums[low]
//}
func findMin(nums []int) int {
	low, high := 0, len(nums)-1
	for low < high {
		pivot := low + (high-low)/2
		if nums[pivot] < nums[high] {
			high = pivot
		} else {
			low = pivot + 1
		}
	}
	return nums[low]
}

//
//func findMin(nums []int) int {
//	length := len(nums)
//	index1 := 0
//	index2 := length-1
//	indexMid := index1
//
//	for nums[index1] >= nums[index2]{
//		if index2 - index1 == 1 {
//			indexMid = index2
//			break
//		}
//		indexMid = (index1 + index2) / 2
//		if nums[indexMid] >= nums[index1]{
//			index1 = indexMid
//		} else if nums[indexMid] <= nums[index2]{
//			index2 = indexMid
//		}
//	}
//	return nums[indexMid]
//}

func main() {
	nums := []int{3, 4, 8, 1, 2}
	res := findMin(nums)
	fmt.Println(res)
}
