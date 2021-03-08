package main

import "fmt"

func mincostTickets(days []int, costs []int) int {
	// 日期数组藏长度，days的最后一天就是出行的天数，下标从0开始 所以+1
	dpCount := days[len(days)-1] + 1

	// DP数组
	dpArray := make([]int, dpCount)

	// DP 出行日 初始化
	for _, day := range days {
		dpArray[day]++
	}

	// 遍历DP数组
	for i := 1; i < dpCount; i++ {
		// 对DP标记的出行日进行动态规划的计算
		// DP状态转移
		if dpArray[i] > 0 {
			// 确定DP关系式
			dpArray[i] = min(
				dpArray[i-1]+costs[0],
				dpArray[max(i-7, 0)]+costs[1],
				dpArray[max(i-30, 0)]+costs[2],
			)
		} else {
			// 如果今天不出行，则今天的花费还是昨天的花费
			dpArray[i] = dpArray[i-1]
		}
	}
	return dpArray[dpCount-1]
}

func main() {
	costs := []int{2, 7, 12}
	days := []int{1, 4, 6, 7, 8, 20}
	a := mincostTickets(days, costs)
	fmt.Println(a)
}

func min(a, b, c int) int {
	//result := 0
	if a < b {
		b = a
	} else {
		a = b
	}
	if a < c {
		c = a
	}
	return c
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}
