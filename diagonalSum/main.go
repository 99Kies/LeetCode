package main

func diagonalSum(mat [][]int) int {
	col_len := len(mat[0])
	sum := 0
	for i := range mat {
		for j := range mat[i] {
			if i == j {
				sum = sum + mat[i][j]
			}

		}
		sum = sum + mat[i][col_len-i-1]
	}
	flag := 0
	if col_len%2 == 1 {
		flag = mat[col_len/2][col_len/2]
	}
	return sum - flag
}

func main() {
	mat := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	diagonalSum(mat)
}
