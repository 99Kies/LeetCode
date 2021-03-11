package main

import "fmt"

func quickSort(nums []int, left int, right int) {
	if left >= right {
		return
	}

	i, j := left, right
	p := nums[i]
	for i < j {
		for i < j && p < nums[j] {
			j--
		}
		nums[i] = nums[j]
		for i < j && p > nums[i] {
			i++
		}
		nums[i] = nums[j]
	}
	nums[j] = p
	//quickSort(nums, left, i-1)
	//quickSort(nums, i+1, right)
}

func quickSortmy(nums []int, left int, right int) {
	if left >= right {
		return
	}
	i, j := left, right
	p := nums[i]
	for i < j {
		// 升序
		for i < j && nums[i] < p {
			// 从左边遍历 找到一个大于基准数的数值， 大的放右边
			i++
		}
		for i < j && nums[j] > p {
			// 从右边遍历 找到一个小于基准数的数值， 小的放左边
			j--
		}
		nums[i], nums[j] = nums[j], nums[i]
	}
	nums[j] = p

	quickSortmy(nums, left, i-1)
	quickSortmy(nums, i+1, right)
}

func quickSortMy2(nums []int, left int, right int) {
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
	quickSortMy2(nums, left, i-1)
	quickSortMy2(nums, i+1, right)
}

func quickSortMy3(nums []int, left int, right int) {
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
	quickSortMy3(nums, left, i-1)
	quickSortMy3(nums, i+1, right)

}

func sortArray(nums []int) []int {
	fmt.Println(nums)
	fmt.Println(len(nums) - 1)
	quickSortLeetCode(nums, 0, len(nums)-1)
	fmt.Println(nums)
	return nums
}
func quickSortLeetCode(nums []int, left int, right int) {
	if left >= right {
		return
	}
	i, j := left, right
	p := nums[i]
	fmt.Println(p)
	for i < j {
		fmt.Println(p)
		for i < j && nums[i] < p {
			i++
		}
		for i < j && nums[j] > p {
			j--
		}
		nums[i], nums[j] = nums[j], nums[i]
	}
	nums[j] = p
	quickSortLeetCode(nums, left, i-1)
	quickSortLeetCode(nums, i+1, right)
}

func main() {
	a := []int{6, 5, 3, 1, 8, 7, 2, 4}
	//quickSortMy2(a, 0, len(a)-1)
	//fmt.Println(a)
	quickSorthuifeng(a, 0, len(a)-1)
	fmt.Println(a)
}

func quickSorthuifeng(nums []int, left, right int) {
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
	quickSorthuifeng(nums, left, i-1)
	quickSorthuifeng(nums, i+1, right)
}

func quickgiaogiao(nums []int, left, right int) {
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
	quickgiaogiao(nums, left, i-1)
	quickgiaogiao(nums, i+1, right)
}
