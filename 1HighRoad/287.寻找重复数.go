package main

import "fmt"

/*
 * @lc app=leetcode.cn id=287 lang=golang
 *
 * [287] 寻找重复数
 */

// @lc code=start
func findDuplicate(nums []int) int {
	mp := map[int]bool{}

	for _, key := range nums {
		if _, ok := mp[key]; ok {
			return key
		}
		mp[key] = true
	}
	return -1

}

// @lc code=end

func main() {
	nums := []int{1, 3, 4, 2, 2}
	res := findDuplicate(nums)
	fmt.Println(res)
}
