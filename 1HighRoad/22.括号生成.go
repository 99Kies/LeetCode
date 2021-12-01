package main

import "fmt"

/*
 * @lc app=leetcode.cn id=22 lang=golang
 *
 * [22] 括号生成
 */

// var ResI = []string{}

// @lc code=start
func generateParenthesis(n int) []string {
	var resI = &[]string{}
	if n <= 0 {
		return *resI
	}

	dfs(n, "", resI, 0, 0)
	return *resI
}

func dfs(n int, path string, res *[]string, open int, close int) {
	if len(path) == 2*n {
		*res = append(*res, path)
		return
	}

	if open < n {
		dfs(n, path+"(", res, open+1, close)
	}
	if close < open {
		dfs(n, path+")", res, open, close+1)
	}
}

// @lc code=end

func main() {
	a := generateParenthesis(3)
	fmt.Println(a)
}
