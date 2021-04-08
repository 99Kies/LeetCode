package main

import "fmt"

func main() {
	n := 0
	num := 0
	fmt.Scanln(&n)
	fmt.Scanln(&num)
	var nums []int
	for num > 0 {
		nums = append(nums, num%10)
		num = num / 10
	}

	flag := -1
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if j == i {
				continue
			}
			if nums[j]+nums[i] == 10 {
				flag = nums[j]
				nums = append(nums[:j], nums[j+1:]...)
				n = n - 1
			}
			if i > 0 {
				if nums[i-1]+flag == 10 {
					nums = append(nums[:i-1], nums[i:]...)
					flag = -1
					n = n - 1
				}
			} else if i == 0 {
				if nums[i]+flag == 10 {
					nums = append(nums[:i], nums[i+1:]...)
					flag = -1
					n = n - 1
				}
			}
		}
	}
	length := len(nums)
	fmt.Println(length)
}
