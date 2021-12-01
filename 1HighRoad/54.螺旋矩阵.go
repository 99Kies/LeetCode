package main

/*
 * @lc app=leetcode.cn id=54 lang=golang
 *
 * [54] 螺旋矩阵
 */

// @lc code=start
func spiralOrder(matrix [][]int) []int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return []int{}
	}

	rows, columns := len(matrix), len(matrix[0])
	order := make([]int, rows*columns)
	index := 0
	left, right, top, bottom := 0, columns-1, 0, rows-1

	for left <= right && top <= bottom {
		for colum := left; colum <= right; colum++ {
			order[index] = matrix[top][colum]
			index++
		}
		for row := top + 1; row <= bottom; row++ {
			order[index] = matrix[row][right]
			index++
		}
		if left < right && top < bottom {
			for column := right - 1; column > left; column-- {
				order[index] = matrix[bottom][column]
				index++
			}
			for row := bottom; row > top; row-- {
				order[index] = matrix[row][left]
				index++
			}
		}
		left++
		right--
		top++
		bottom--
	}
	return order
}

// @lc code=end

func main() {

}
