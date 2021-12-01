package main

/*
 * @lc app=leetcode.cn id=162 lang=golang
 *
 * [162] 寻找峰值
 */
import "fmt"

// @lc code=start
func findPeakElement(nums []int) (idx int) {
	l, r := 0, len(nums)-1
	for l < r {
		mid := (l + r) / 2
		if nums[mid] < nums[mid+1] {
			l = mid + 1
		} else {
			r = mid
		}
	}
	return l
}

// @lc code=end

func main() {
	a := []int{1, 2, 1, 3, 5, 6, 4}
	fmt.Println(findPeakElement(a))
}
