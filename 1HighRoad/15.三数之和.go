package main

import "sort"

/*
 * @lc app=leetcode.cn id=15 lang=golang
 *
 * [15] 三数之和
 */

// @lc code=start
func threeSum(nums []int) [][]int {
	n := len(nums)
	if nums == nil || n < 3 {
		return [][]int{}
	}
	sort.Ints(nums)
	res := [][]int{}
	for i := 0; i < n; i++ {
		if nums[i] > 0 {
			return res
		}
		if i > 0 && nums[i] == nums[i-1] {
			// 跳过重复的值，防止获得重复的解。
			continue
		}
		L := i + 1
		R := n - 1
		for L < R {
			if nums[i]+nums[L]+nums[R] == 0 {
				res = append(res, []int{nums[i], nums[L], nums[R]})
				for L < R && nums[L] == nums[L+1] {
					// 跳过重复的值，防止获得重复的解。
					L = L + 1
				}
				for L < R && nums[R] == nums[R-1] {
					// 跳过重复的值，防止获得重复的解。
					R = R - 1
				}
				L = L + 1
				R = R - 1
			} else if nums[i]+nums[L]+nums[R] > 0 {
				R = R - 1
			} else {
				L = L + 1
			}
		}
	}
	return res
}

// @lc code=end

func main() {

}
