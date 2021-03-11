package main

import (
	"fmt"
)

func selectSort(nums []int) []int {
	length := len(nums)

	for i := 0; i < length; i++ {
		maxIndex := 0
		for j := 1; j < length-i; j++ {
			if nums[j] > nums[maxIndex] {
				maxIndex = j
			}
		}
		nums[length-i-1], nums[maxIndex] = nums[maxIndex], nums[length-i-1]
	}
	fmt.Println(nums)
	return nums
}

func selectSortMy(nums []int) []int {
	length := len(nums)

	for i := 0; i < length; i++ {
		k := i // 以i作为下标
		for j := i + 1; j < length; j++ {
			// 寻找最大值 然后用k标记
			if nums[j] > nums[k] {
				k = j
			}
		}
		if k != i { // 如果找到了一个最大值就和i替换
			nums[k], nums[i] = nums[i], nums[k]
		}

		fmt.Println(nums)
	}
	return nums
}

func SelectSort(data []int) {
	length := len(data)
	for i := 0; i < length; i++ {
		tmp := data[i]
		flag := i
		for j := i + 1; j < length; j++ {
			if data[j] < tmp {
				tmp = data[j]
				flag = j
			}
		}

		if flag != i {
			data[flag] = data[i]
			data[i] = tmp
		}
		fmt.Println(data) //为了看具体排序的过程

	}
}

func selectSortmy2(nums []int) []int {
	k := 0
	length := len(nums)
	for i := 0; i < length; i++ {
		k = i
		for j := i; j < length; j++ {
			if nums[j] < nums[k] {
				k = j
			}
		}
		if k != i {
			nums[k], nums[i] = nums[i], nums[k]

		}
	}
	fmt.Println(nums)
	return nums
}

func selectSortMy3(nums []int) []int {
	length := len(nums)
	for i := 0; i < length-1; i++ {
		k := i
		for j := i + 1; j < length; j++ {
			if nums[j] < nums[k] {
				k = j
			}
		}
		if k != i {
			nums[i], nums[k] = nums[k], nums[i]
		}
	}
	return nums
}

func select_(nums []int) []int {
	length := len(nums)
	for i := 0; i < length-1; i++ {
		k := i
		for j := i; j < length; j++ {
			if nums[k] > nums[j] {
				k = j
			}
		}
		if k != i {
			nums[k], nums[i] = nums[i], nums[k]
		}
	}
	return nums
}

func main() {
	a := []int{1, 2, 5, 3, 4, 7, 6}
	select_(a)
	fmt.Println(a)
}
