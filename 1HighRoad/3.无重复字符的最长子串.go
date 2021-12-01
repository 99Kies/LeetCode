package main

func lengthOfLongestSubstring(s string) int {
	m := map[byte]int{}
	n := len(s)

	rk, ans := -1, 0
	for i := 0; i < n; i++ {
		var maxLen int
		if i != 0 {
			// 重置上一个开头的值
			delete(m, s[i-1])
		}
		for rk+1 < n && m[s[rk+1]] == 0 {
			m[s[rk+1]]++
			rk++
			maxLen++
		}
		ans = max(ans, maxLen)
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
