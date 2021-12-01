package main

import (
	"fmt"
	"math"
	"strconv"
)

/*
 * @lc app=leetcode.cn id=556 lang=golang
 *
 * [556] 下一个更大元素 III
 */

// @lc code=start
func nextGreaterElement(n int) int {
	str := []byte(strconv.Itoa(n))

	// fmt.Println(str)
	// str = reverse(str, len(str)/2)
	// fmt.Println(str)
	// return 0
	i := len(str) - 2
	for i >= 0 && str[i+1] <= str[i] {
		i--
	}
	if i < 0 {
		return -1
	}
	j := len(str) - 1
	for j >= 0 && str[i] >= str[j] {
		j--
	}
	str[i], str[j] = str[j], str[i]
	fmt.Println(i, j, len(str))
	fmt.Println(strconv.Atoi(string(str)))
	reverse(str, i+1)
	fmt.Println(strconv.Atoi(string(str)))
	res, _ := strconv.Atoi(string(str))
	fmt.Println(res)

	if res > math.MaxInt32 {
		return -1
	}
	return res
}

func reverse(str []byte, start int) {
	i, j := start, len(str)-1
	for i < j {
		str[i], str[j] = str[j], str[i]
		i++
		j--
	}
}

// @lc code=end

func main() {
	a := nextGreaterElement(2147483476)

	fmt.Println(a)
	fmt.Println(math.MaxInt32)
	fmt.Println(1 << 31)
}

// 2147483684
// 2147483486
