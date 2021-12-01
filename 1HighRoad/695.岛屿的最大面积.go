package main

/*
 * @lc app=leetcode.cn id=695 lang=golang
 *
 * [695] 岛屿的最大面积
 */

// @lc code=start
func DFSLand(grid [][]int, i, j int) int {
	if i < len(grid) && i >= 0 && j < len(grid[0]) && j >= 0 {
		if grid[i][j] == 0 {
			return 0
		} else {
			grid[i][j] = 0
			return 1 + DFSLand(grid, i-1, j) + DFSLand(grid, i+1, j) + DFSLand(grid, i, j-1) + DFSLand(grid, i, j+1)
		}
	} else {
		return 0
	}
}

func maxAreaOfIsland(grid [][]int) int {
	maxRes := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			maxRes = max(maxRes, DFSLand(grid, i, j))
		}
	}
	return maxRes
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end

func main() {

}
