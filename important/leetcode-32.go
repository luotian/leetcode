package important

//32. 下一个排列
//给你一个只包含 '(' 和 ')' 的字符串，找出最长有效（格式正确且连续）括号子串的长度。
//
//示例 1：
//输入：s = "(()"
//输出：2
//解释：最长有效括号子串是 "()"
//示例 2：
//输入：s = ")()())"
//输出：4
//解释：最长有效括号子串是 "()()"
//示例 3：
//输入：s = ""
//输出：0
//
//提示：
//0 <= s.length <= 3 * 104
//s[i] 为 '(' 或 ')'

//题解：使用栈的思想方法
// 还有动态规划的方式，还未入门
func longestValidParentheses(s string) int {
	if len(s) < 2 {
		return 0
	}

	var maxLen int
	var stack []int = []int{-1}
	for i, v := range s {
		if v == '(' {
			stack = append(stack, i)
		} else if v == ')' {
			stack = stack[:len(stack)-1]
			if len(stack) > 0 {
				l := i - stack[len(stack)-1]
				if l > maxLen {
					maxLen = l
				}
			}
			if len(stack) == 0 {
				stack = append(stack, i)
			}
		} else {
			panic("xxx")
		}
	}

	return maxLen
}
