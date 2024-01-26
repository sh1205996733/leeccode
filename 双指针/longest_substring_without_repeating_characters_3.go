package pointer

// 无重复字符的最长子串
//https://leetcode.cn/problems/longest-substring-without-repeating-characters/submissions/66227377/

// 滑动窗口 同向双指针
func lengthOfLongestSubstring0(s string) int {
	ret := 0
	left := 0
	// 保存遍历过的字符的最新下标
	store := make(map[uint8]int)
	for i := 0; i < len(s); i++ {
		if ci, ok := store[s[i]]; ok { // 如果i字符之前出现过
			left = max(left, ci)
		}
		ret = max(ret, i-left+1)
		// 存储下标时，以1为起始位置，为了保证上面的 left = max(left, ci) 的正确性
		store[s[i]] = i + 1
	}

	return ret
}

// 滑动窗口 同向双指针
// O(n)
// O(n)
func lengthOfLongestSubstring(s string) int {
	ans := 0
	left := 0
	// 保存遍历过的字符的最新下标
	store := make(map[uint8]int)
	for right := 0; right < len(s); right++ {
		store[s[right]]++
		for store[s[right]] > 1 { // 如果i字符之前出现过
			store[s[left]]--
			left++
		}
		ans = max(ans, right-left+1)
	}
	return ans
}
