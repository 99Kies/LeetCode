package main

/*
 * @lc app=leetcode.cn id=152 lang=golang
 *
 * [152] 乘积最大子数组
 */

// @lc code=start
func maxProduct(nums []int) int {
	max := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i-1] > 0 {
			// 如果前一个数大于0 那当前数就加上前一个数。
			nums[i] *= nums[i-1]
		}

		if max < nums[i] {
			// 记录最大值
			max = nums[i]
		}
	}
	return max

}

// @lc code=end

func main() {

}
