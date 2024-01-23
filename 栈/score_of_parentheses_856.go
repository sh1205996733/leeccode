package stack

// 括号的分数
// https://leetcode.cn/problems/score-of-parentheses/description/
// 栈
func scoreOfParentheses(s string) int {
	stack := []int{0}
	for _, c := range s {
		if c == '(' {
			stack = append(stack, 0)
		} else {
			v := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			stack[len(stack)-1] += max(2*v, 1)
		}
	}
	return stack[0]
}

func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}

func scoreOfParentheses0(s string) (ans int) {
	bal := 0
	for i, c := range s {
		if c == '(' {
			bal++
		} else {
			bal--
			if s[i-1] == '(' {
				ans += 1 << bal
			}
		}
	}
	return
}
