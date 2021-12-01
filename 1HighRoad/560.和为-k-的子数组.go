package main

/*
 * @lc app=leetcode.cn id=560 lang=golang
 *
 * [560] 和为 K 的子数组
 */

// @lc code=start
func subarraySum(nums []int, k int) int {
	count, pre := 0, 0
	mp := map[int]int{}
	mp[0] = 1
	for i := 0; i < len(nums); i++ {
		pre += nums[i] // 计算前序数组
		if _, ok := mp[pre-k]; ok {
			count += mp[pre-k]
		}
		mp[pre] += 1
	}
	return count
}

// @lc code=end

func main() {

}
