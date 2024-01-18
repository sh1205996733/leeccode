package string

// 最长公共前缀
// https://leetcode.cn/problems/longest-common-prefix/description/
func longestCommonPrefix0(strs []string) string {
	count := len(strs)
	min := 0
	for i := 0; i < count; i++ {
		if len([]byte(strs[min])) > len([]byte(strs[i])) {
			min = i
		}
	}
	index := 0
	prefix := strs[min]
	for index < len([]byte(prefix)) { //strs元素不为空
		c := string(prefix[index])
		for i := 0; i < count; i++ {
			if c != string(strs[i][index]) {
				return prefix[0:index]
			}
		}
		index++
	}
	return prefix
}

// 纵向扫描
func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	for i := 0; i < len(strs[0]); i++ {
		for j := 1; j < len(strs); j++ {
			if i == len(strs[j]) || strs[j][i] != strs[0][i] {
				return strs[0][:i]
			}
		}
	}
	return strs[0]
}
