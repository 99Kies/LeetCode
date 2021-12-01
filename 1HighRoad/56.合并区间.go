package main

import "sort"

/*
 * @lc app=leetcode.cn id=56 lang=golang
 *
 * [56] 合并区间
 */

// @lc code=start
func merge(intervals [][]int) [][]int {
	if len(intervals) <= 1 {
		return intervals
	}
	sort.Slice(intervals, func(i, j int) bool {
		// 按照第一个元素进行排序
		return intervals[i][0] < intervals[j][0]
	})

	x, y := intervals[0][0], intervals[0][1]
	result := make([][]int, 0)
	for i := 1; i < len(intervals); i++ {
		if y >= intervals[i][0] {
			if y < intervals[i][1] {
				y = intervals[i][1]
			}
		} else {
			result = append(result, []int{x, y})
			x = intervals[i][0]
			y = intervals[i][1]
		}
	}
	result = append(result, []int{x, y})

	return result
}

// @lc code=end

func main() {

}
