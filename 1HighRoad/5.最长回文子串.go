/*
 * @lc app=leetcode.cn id=5 lang=golang
 *
 * [5] 最长回文子串
 */

// @lc code=start

func longestPalindrome(s string) string {
	start, end := 0, 0
	for i := range s {
		l, r := expand(s, i, i) // 奇数的情况
		if end-start < r-l {    // 如果有更大的边界值
			start, end = l, r
		}

		l, r = expand(s, i, i+1) // 偶数的情况
		if end-start < r-l {     // 如果有更大的边界值
			start, end = l, r
		}
	}
	return s[start : end+1]
}

// 围绕中心展开
func expand(s string, left, right int) (int, int) {
	// 这个是核心,给一个字符串+字符串左右边界,找出该范围内最长的回文字符串的左右边界

	for left >= 0 && right < len(s) && s[left] == s[right] {
		left, right = left-1, right+1
	}
	return left + 1, right - 1 // 上面的循环最后那块把多操作的加减回去
}

// @lc code=end

