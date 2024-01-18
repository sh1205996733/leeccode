package monotonicstack

import "container/list"

// 有效的括号
// https://leetcode.cn/problems/valid-parentheses/description/

var map1 = map[rune]int{
	'(': -1,
	')': 1,
	'[': -2,
	']': 2,
	'{': -3,
	'}': 3,
}

// ()[]{}  (] (((())))[[{}]]
func isValid0(s string) bool {
	stack := list.New()
	for _, c := range s {
		if map1[c] < 0 {
			stack.PushBack(c)
		} else if stack.Len() > 0 && map1[stack.Back().Value.(rune)]+map1[c] == 0 {
			stack.Remove(stack.Back())
		} else {
			return false
		}
	}
	return stack.Len() == 0
}

var map2 = map[rune]rune{
	'(': ')',
	'[': ']',
	'{': '}',
}

func isValid(s string) bool {
	stack := make([]rune, 0)
	for _, c := range s {
		if _, ok := map2[c]; ok {
			stack = append(stack, c)
		} else if len(stack) > 0 && map2[stack[len(stack)-1]] == c {
			stack = stack[:len(stack)-1]
		} else {
			return false
		}
	}
	return len(stack) == 0
}
