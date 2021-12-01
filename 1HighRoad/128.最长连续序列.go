/*
 * @lc app=leetcode.cn id=128 lang=golang
 *
 * [128] 最长连续序列
 */

// @lc code=start
func longestConsecutive(nums []int) int {
	numSet := map[int]bool{}
	for _, num := range nums {
		numSet[num] = true
	}
	maxStreak := 0
	// 尝试的每一种长度，取其最大值。
	for num := range numSet {
		if !numSet[num-1] {
			// 如果没有比当前小的数 那就去尝试遍历
			currentNum := num
			currentStreak := 1
			for numSet[currentNum+1] {
				// 如果前面有目标值 就继续走
				currentNum++
				currentStreak++
			}
			if maxStreak < currentStreak {
				maxStreak = currentStreak
			}
		}
	}
	return maxStreak
}

// @lc code=end

