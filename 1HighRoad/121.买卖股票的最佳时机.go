package main

import "fmt"

/*
 * @lc app=leetcode.cn id=121 lang=golang
 *
 * [121] 买卖股票的最佳时机
 */

// @lc code=start
func maxProfit(prices []int) int {
	minPrice := 9999
	maxProfit := 0
	for i := 0; i < len(prices); i++ {
		// 动态规划
		maxProfit = max(prices[i]-minPrice, maxProfit)
		minPrice = min(prices[i], minPrice)
	}
	return maxProfit
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

// @lc code=end

func main() {
	fmt.Println(maxProfit([]int{2, 4, 1}))
}
