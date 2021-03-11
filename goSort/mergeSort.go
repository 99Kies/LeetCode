package main

import (
	"fmt"
	"sort"
)

func main() {
	arr := []int{8, 9, 5, 7, 1, 2, 5, 7, 6, 3, 5, 4, 8, 1, 8, 5, 3, 5, 8, 4}
	result := mergeSortgiao(arr)
	fmt.Println(result)
}

/**
归并排序（Merge sort，台湾译作：合并排序）是建立在归并操作上的一种有效的排序算法。该算法是采用分治法（Divide and Conquer）的一个非常典型的应用
分治思想，时间复杂度为：O(n*log(n))
*/

func mergeSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}
	i := len(arr) / 2
	left := mergeSort(arr[0:i])
	right := mergeSort(arr[i:])
	result := merge(left, right)
	return result
}

func mergeSortMy(nums []int) []int {
	if len(nums) < 2 {
		// 一直拆分 直到拆分为单个元素
		return nums
	}
	i := len(nums) / 2
	left := mergeSortMy(nums[0:i])
	right := mergeSortMy(nums[i:])
	result := mergeMy(left, right)
	return result
}

func mergesort(nums []int) []int {
	if len(nums) < 2 {
		return nums
	}
	i := len(nums) / 2
	left := mergesort(nums[0:i])
	right := mergesort(nums[i:])
	result := mergeMy2(left, right)
	return result
}

func mergeMy2(left, right []int) []int {
	result := make([]int, 0)
	i, j := 0, 0
	l, r := len(left), len(right)
	for i < l && j < r {
		if left[i] > right[j] {
			result = append(result, right[j])
			j++
		} else {
			result = append(result, left[i])
			i++
		}
	}
	result = append(result, right[j:]...)
	result = append(result, left[i:]...)
	return result
}

func mergeSortgiao(nums []int) []int {
	if len(nums) < 2 {
		return nums
	}
	i := len(nums) / 2
	left := mergeSortgiao(nums[0:i])
	right := mergeSortgiao(nums[i:])
	result := mergegiao(left, right)
	return result
}

func mergegiao(left, right []int) []int {
	result := make([]int, 0)
	i, j := 0, 0
	l, r := len(left), len(right)
	for i < l && j < r {
		if left[i] > right[j] {
			result = append(result, right[j])
			j++
		} else {
			result = append(result, left[i])
			i++
		}
	}
	result = append(result, right[j:]...)
	result = append(result, left[i:]...)
	return result
}

func mergeMy(left []int, right []int) []int {
	result := make([]int, 0)
	i, j := 0, 0
	l, r := len(left), len(right)
	for i < l && j < r {
		if left[i] > right[j] {
			result = append(result, right[j])
			j++
		} else {
			result = append(result, left[i])
			i++
		}
	}
	result = append(result, right[j:]...)
	result = append(result, left[i:]...)
	return result
}

func mergeSorthuifeng(nums []int) []int {
	if len(nums) < 2 {
		return nums
	}
	i := len(nums) / 2
	left := mergeSorthuifeng(nums[0:i])
	right := mergeSorthuifeng(nums[i:])
	result := mergehuifeng(left, right)
	return result
}

func mergehuifeng(left, right []int) []int {
	result := make([]int, 0)
	i, j := 0, 0
	l, r := len(left), len(right)
	for i < l && j < r {
		if left[i] > right[j] {
			result = append(result, right[j])
			j++
		} else {
			result = append(result, left[i])
			i++
		}
	}
	result = append(result, right[j:]...)
	result = append(result, left[i:]...)
	return result
}

func merge(left, right []int) []int {
	result := make([]int, 0)
	m, n := 0, 0 // left和right的index位置
	l, r := len(left), len(right)
	for m < l && n < r {
		if left[m] > right[n] {
			result = append(result, right[n])
			n++
			continue
		}
		result = append(result, left[m])
		m++
	}
	result = append(result, right[n:]...) // 这里竟然没有报数组越界的异常？
	result = append(result, left[m:]...)
	return result
}

func mergeCopy(left, right []int) []int {
	result := make([]int, 0)
	m, n := 0, 0
	l, r := len(left), len(right)
	for m < l && n < r {
		// 比两个数，小的先放在前面
		if left[m] > right[n] {
			// 如果right那侧是小的 那就放在前面
			fmt.Println(result)
			result = append(result, right[n])
			n++
			continue
		}
		// 遍历完一遍右侧的 就把左侧比较好的数据放在屁股后面
		result = append(result, left[m])
		m++
	}
	fmt.Println("==============================")
	fmt.Println(left, right)
	fmt.Println(left[m:], right[n:])
	fmt.Println(result)
	fmt.Println()
	// 合并通过筛选的数据，根据比较顺序，所以right那侧是大的，
	result = append(result, right[n:]...)
	result = append(result, left[m:]...)
	return result
}

func mergeSortInt(left, right []int) []int {
	copy(left[:], right[:])
	sort.Ints(left)
	return left
}
