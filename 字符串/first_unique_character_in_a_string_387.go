package string

import "fmt"

// 字符串中的第一个唯一字符
// https://leetcode.cn/problems/first-unique-character-in-a-string/

// 哈希
func firstUniqChar(str string) int {
	var ans = -1
	//m := make(map[uint8]int)
	m := make([]uint8, 26)
	for i := range str {
		m[str[i]-'a']++
	}
	fmt.Println(m)
	for i := range str {
		if m[str[i]-'a'] == 1 {
			ans = i
			break
		}
	}
	return ans
}
