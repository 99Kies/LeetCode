package main

import "fmt"

func majorityElement(nums []int) int {
	length := len(nums)
	quickSort(nums, 0, length-1)
	fmt.Println(nums)
	return 0
}

func selectSort(nums []int) int {
	length := len(nums)
	for i := 0; i < length-1; i++ {
		k := i
		for j := i; j < length; j++ {
			if nums[k] > nums[j] {
				k = j
			}
		}
		if k != i {
			nums[i], nums[k] = nums[k], nums[i]
		}
	}
	return nums[length/2]
}

func quickSort(nums []int, left, right int) {
	if left >= right {
		return
	}
	i, j := left, right
	p := nums[i]
	for i < j {
		for i < j && nums[i] < p {
			i++
		}
		for i < j && nums[j] > p {
			j--
		}
		nums[i], nums[j] = nums[j], nums[i]
	}
	nums[j] = p
	quickSort(nums, left, i-1)
	quickSort(nums, i+1, right)
}

func main() {
	a := []int{2, 2, 1, 1, 1, 2, 2}
	fmt.Println(a)
	b := selectSort(a)
	fmt.Println(b)
}
