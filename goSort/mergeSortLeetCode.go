package main

import "fmt"

func main() {
	a := []int{2, 3, 4, 1, 8, 5}
	fmt.Println(sortArrayT(a))
}

func sortArrayT(nums []int) []int {
	if len(nums) <= 1 {
		return nums
	}

	return mergeSortT(nums, 0, len(nums)-1)
}

func mergeSortT(nums []int, left, right int) []int {
	if left == right {
		return []int{nums[left]}
	}

	mid := (left + right) / 2
	ret1 := mergeSortT(nums, left, mid)
	ret2 := mergeSortT(nums, mid+1, right)
	return mergeT(ret1, ret2)
}

func mergeT(ret1 []int, ret2 []int) []int {
	var ret []int

	i, j := 0, 0
	for i < len(ret1) && j < len(ret2) {
		if ret1[i] < ret2[j] {
			ret = append(ret, ret1[i])
			i++
		} else {
			ret = append(ret, ret2[j])
			j++
		}
	}

	for ; i < len(ret1); i++ {
		ret = append(ret, ret1[i])
	}
	for ; j < len(ret2); j++ {
		ret = append(ret, ret2[j])
	}

	return ret
}
