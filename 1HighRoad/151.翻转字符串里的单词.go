package main

import (
	"fmt"
)

/*
 * @lc app=leetcode.cn id=151 lang=golang
 *
 * [151] 翻转字符串里的单词
 */

// @lc code=start
func reverseWords(s string) string {
	b := []byte(s)
	// 先把整体进行一次旋转
	reverse(b, 0, len(b)-1)
	start := 0
	end := 0
	i := 0
	res := ""
	for i < len(b) {
		if b[i] != ' ' {
			start = i
			for i < len(b) && b[i] != ' ' {
				i += 1
			}
			end = i - 1
			// 随后把字符串里的单个单词进行旋转（这样子单词就“正”了）
			reverse(b, start, end)
			res = res + " " + string(b[start:end+1])
		}
		i += 1
	}
	return res[1:]
}

func reverse(b []byte, start, end int) {
	// 翻转 byte 数组
	for start < end {
		b[start], b[end] = b[end], b[start]
		start++
		end--
	}
}

// @lc code=end

func main() {
	a := reverseWords("  hello world my name is bob   ")
	fmt.Println(a)
}
