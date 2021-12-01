/*
 * @lc app=leetcode.cn id=11 lang=golang
 *
 * [11] 盛最多水的容器
 */
package main

import "fmt"

// @lc code=start
func maxArea(height []int) int {
	len := len(height)
	i, j := 0, len-1

	_maxArea := min(height[i], height[j]) * (j - i)
	if min(height[i], height[j]) == height[i] {
		i++
	} else {
		j--
	}
	area := 0
	for i <= j {
		if height[i] == min(height[i], height[j]) {
			area = height[i] * (j - i)
			i++
		} else {
			area = height[j] * (j - i)
			j--
		}
		if area > _maxArea {
			_maxArea = area
		}
	}
	return _maxArea
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

// @lc code=end

func main() {
	a := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	res := maxArea(a)
	fmt.Println(res)
}
