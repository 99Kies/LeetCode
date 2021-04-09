// dp, 动态规划, 使用最小花费爬楼梯

package main

import "fmt"

func minCostClimbingStairs(cost []int) int {
	length := len(cost)
	dp := make([]int, length+1)
	for i := 2; i <= length; i++ {
		dp[i] = min(dp[i-1]+cost[i-1], dp[i-2]+cost[i-2])
		fmt.Println(dp, dp[i-1]+cost[i-1], dp[i-2]+cost[i-2])
	}

	return dp[length]

}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func main() {
	cost := []int{10, 15, 20}
	res := minCostClimbingStairs(cost)
	fmt.Println(res)
}
