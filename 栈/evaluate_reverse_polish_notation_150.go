package stack

import "strconv"

// 逆波兰表达式求值
// https://leetcode.cn/problems/evaluate-reverse-polish-notation/description/

// 栈
func evalRPN(tokens []string) int {
	stack := make([]int, 0) //储存中间结果的栈
	for _, ss := range tokens {
		if isNumber(ss) { // 操作数
			a, _ := strconv.Atoi(ss)
			stack = append(stack, a)
		} else {
			b := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			a := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			result := 0
			if ss == "+" {
				result = a + b
			} else if ss == "-" {
				result = a - b
			} else if ss == "*" {
				result = a * b
			} else if ss == "/" {
				result = a / b
			}
			stack = append(stack, result)
		}
	}
	return stack[len(stack)-1]
}
func isNumber(token string) bool {
	return !("+" == token || "-" == token || "*" == token || "/" == token)
}
