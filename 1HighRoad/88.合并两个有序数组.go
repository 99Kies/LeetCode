/*
 * @lc app=leetcode.cn id=88 lang=golang
 *
 * [88] 合并两个有序数组
 */

// @lc code=start
func merge(nums1 []int, m int, nums2 []int, n int) {
	p1, p2, tail := m-1, n-1, m+n-1
	for p1 >= 0 || p2 >= 0 {
		var cur int
		// cur 记录最大值
		if p1 == -1 {
			// 边缘条件，如果遍历到nums1开头之外，则p2还剩下一个数还未存储
			cur = nums2[p2]
			p2--
		} else if p2 == -1 {
			// 反之
			cur = nums1[p1]
			p1--
		} else if nums1[p1] > nums2[p2] {
			cur = nums1[p1]
			p1--
		} else {
			cur = nums2[p2]
			p2--
		}
		nums1[tail] = cur
		tail--
	}
}

// @lc code=end

