package string

import "strings"

// 验证回文串
//https://leetcode.cn/problems/valid-palindrome/description/

func isPalindrome0(str string) bool {
	str = strings.ToLower(str)
	length := len(str)
	head, tail := 0, length-1
	for head < tail {
		if !isVaild(str[head]) {
			head++
		} else if !isVaild(str[tail]) {
			tail--
		} else if str[head] == str[tail] {
			head++
			tail--
		} else {
			return false
		}
	}
	return true
}
func isPalindrome(str string) bool {
	tmp := make([]byte, len(str))
	index := 0
	for i := 0; i < len(str); i++ {
		if isVaild(str[i]) {
			if str[i] >= 'A' && str[i] <= 'Z' {
				tmp[index] = str[i] + 32
			} else {
				tmp[index] = str[i]
			}
			index++
		}
	}
	head, tail := 0, index-1
	for head < tail {
		if tmp[head] == tmp[tail] {
			head++
			tail--
		} else {
			return false
		}
	}
	return true
}

func isVaild(c byte) bool {
	return (c >= '0' && c <= '9') || (c >= 'a' && c <= 'z')
}
