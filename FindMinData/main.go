// 旋转数组的最小数字
// 查找问题，通过二分查找实现

package main

import "fmt"

func FindMinData(numbers []int, length int) int {
	if numbers == nil || length <= 0 {
		return 0
	}
	index1 := 0
	index2 := length - 1
	indexMid := index1
	for numbers[index1] >= numbers[index2] {
		if index2-index1 == 1 {
			indexMid = index2
			break
		}
		indexMid = (index1 + index2) / 2
		if numbers[indexMid] >= numbers[index1] {
			index1 = indexMid
		} else if numbers[indexMid] <= numbers[index2] {
			index2 = indexMid
		}

	}
	return numbers[indexMid]
}

//func binarySearch(nums []int, target int) {
//	left := 0
//	right := ...
//
//	for ... {
//		mid := (right + left) / 2
//		if nums[mid] == target {
//			...
//		} else if nums[mid] < target {
//			left = ...
//		} else if nums[mid] > target {
//			right = ...
//		}
//	}
//	return ...
//}

func main() {
	a := []int{3, 4, 5, 1, 2}
	min := FindMinData(a, 5)
	fmt.Println(min)
}
