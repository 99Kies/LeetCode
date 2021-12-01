package main

/*
 * @lc app=leetcode.cn id=209 lang=golang
 *
 * [209] 长度最小的子数组
 */

// @lc code=start
func minSubArrayLen(target int, nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}

	start, end := 0, 0
	// 两个指针都从头开始
	minLen := n + 1
	sum := 0
	for end < n {
		sum += nums[end]
		for sum >= target {
			// 如果 sum 比 target 高，则右侧的end就停止移动，让start开始前移。
			mlen := end - start + 1
			if mlen < minLen {
				minLen = mlen
			}
			sum -= nums[start]
			start++
		}
		// 如果该end结尾的start->end区间里的内容没有期望的值了，那就end迁移，继续寻找。
		end++
	}
	if minLen == n+1 {
		return 0
	}
	return minLen
}

// @lc code=end

func main() {

}
