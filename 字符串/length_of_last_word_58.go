package string

// 最后一个单词的长度
// https://leetcode.cn/problems/length-of-last-word/description/

func lengthOfLastWord(s string) (ans int) {
	index := len(s) - 1
	for s[index] == ' ' {
		index--
	}
	for index > -1 && s[index] != ' ' {
		ans++
		index--
	}
	return
}
