/*
 * @lc app=leetcode.cn id=415 lang=golang
 *
 * [415] 字符串相加
 */

// @lc code=start
func addStrings(num1 string, num2 string) string {
	add := 0
	ans := ""

	for i, j := len(num1)-1, len(num2)-1; i >= 0 || j >= 0 || add != 0; i, j = i-1, j-1 {
		var x, y int

		if i >= 0 {
			x = int(num1[i] - '0')
		}

		if j >= 0 {
			y = int(num2[j] - '0')
		}
		result := x + y + add
		// strconv.Itoa(123) 把 整数类型 转成 字符串类型
		ans = strconv.Itoa(result%10) + ans // 把 个位上 的数累加在结果上
		add = result / 10                   // 把 十位上 的数，用于下次计算
	}
	return ans
}

// @lc code=end

