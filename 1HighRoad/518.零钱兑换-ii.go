package main

/*
 * @lc app=leetcode.cn id=518 lang=golang
 *
 * [518] 零钱兑换 II
 */

// @lc code=start
func change(amount int, coins []int) int {
	dp := make([]int, amount+1)
	dp[0] = 1
	for i := 1; i <= len(coins); i++ {
		val := coins[i-1]
		for j := val; j <= amount; j++ {
			dp[j] += dp[j-val]
		}
	}
	return dp[amount]
}

// @lc code=end

func main() {

}
