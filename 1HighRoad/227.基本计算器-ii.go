/*
 * @lc app=leetcode.cn id=227 lang=golang
 *
 * [227] 基本计算器 II
 */

// @lc code=start
func calculate(s string) int {
	stack := []int{}
	preSign := '+'
	num := 0
	for i, ch := range s {
		isDigit := '0' <= ch && ch <= '9'
		if isDigit {
			// 把字符串数字截取出来
			num = num*10 + int(ch-'0')
		}
		if !isDigit && ch != ' ' || i == len(s)-1 {
			// 进行运算，先运算 * 和 /，+ 和 - 的运算先推入到栈内。
			switch preSign {
			// 如果ch为符号，则进行上次的运算。（这样子的目的是为了找到 n * m - x 中的 m，这样子才能实现前面 n * m的操作）
			case '+':
				stack = append(stack, num)
			case '-':
				stack = append(stack, -num)
			case '*':
				stack[len(stack)-1] *= num
			default:
				stack[len(stack)-1] /= num
			}
			preSign = ch // 为下次计算记录符号。
			num = 0
		}
	}
	ans := 0
	for _, v := range stack {
		// 最后进行 + - 的运算。
		ans += v
	}
	return ans
}

// @lc code=end

