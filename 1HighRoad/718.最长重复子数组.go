package main

/*
 * @lc app=leetcode.cn id=718 lang=golang
 *
 * [718] 最长重复子数组
 */

// @lc code=start
func findLength(nums1 []int, nums2 []int) int {
	return
}

func maxLen(a []int, i int, b []int, j, len int) int {
	// 比较两个滑动窗口当中最长的相同区域。
	count, _max := 0, 0
	for k := 0; k < len; k++ {
		if a[i+k] == b[j+k] {
			count++
		} else if count > 0 {
			_max = max(_max, count)
			count = 0
		}
	}
	if count > 0 {
		return max(_max, count)
	}
	return _max
}

// func max(a, b int) int {
// 	if a > b {
// 		return a
// 	}
// 	return b
// }

// @lc code=end

func main() {

}
