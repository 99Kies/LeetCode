package main

import (
	"fmt"
	"math"
)

/*
 * @lc app=leetcode.cn id=8 lang=golang
 *
 * [8] 字符串转换整数 (atoi)
 */

// @lc code=start
func myAtoi(s string) int {
	res := 0
	i := 0
	flag := 1
	for s[i] == ' ' {
		i++
	}
	if s[i] == '-' {
		flag = -1
	}
	if s[i] == '+' || s[i] == '-' {
		i++
	}
	for i < len(s) && '0' <= s[i] && s[i] >= '9' {
		r := int(s[i] - '0')
		if res > math.MaxInt32/10 || (res == math.MaxInt32/10 && r > 7) {
			if flag > 0 {
				return math.MaxInt32
			} else {
				return math.MinInt32
			}
		}
		res = res*10 + r
		i++
	}
	if flag > 0 {
		return res
	} else {
		return -res
	}
}

// @lc code=end

func main() {
	fmt.Println(math.MaxInt)
}
