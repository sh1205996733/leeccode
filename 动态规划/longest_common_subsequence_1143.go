package dp

// 最长公共子序列
// https://leetcode.cn/problems/longest-common-subsequence/solutions/696763/zui-chang-gong-gong-zi-xu-lie-by-leetcod-y7u0/

func longestCommonSubsequence0(text1 string, text2 string) int {
	if text1 == "" || text2 == "" {
		return 0
	}
	chars1 := []rune(text1)
	if len(chars1) == 0 {
		return 0
	}
	chars2 := []rune(text2)
	if len(chars2) == 0 {
		return 0
	}
	rowsChars := chars1
	colsChars := chars2
	if len(chars1) < len(chars2) {
		colsChars = chars1
		rowsChars = chars2
	}
	dp := make([]int, len(colsChars)+1)
	for i := 1; i <= len(rowsChars); i++ {
		cur := 0
		for j := 1; j <= len(colsChars); j++ {
			leftTop := cur
			cur = dp[j]
			if rowsChars[i-1] == colsChars[j-1] {
				dp[j] = leftTop + 1
			} else {
				if dp[j-1] > dp[j] {
					dp[j] = dp[j-1]
				}
			}
		}
	}
	return dp[len(colsChars)]
}

func longestCommonSubsequence(text1, text2 string) int {
	m, n := len(text1), len(text2)
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}
	for i, c1 := range text1 {
		for j, c2 := range text2 {
			if c1 == c2 {
				dp[i+1][j+1] = dp[i][j] + 1
			} else {
				dp[i+1][j+1] = max(dp[i][j+1], dp[i+1][j])
			}
		}
	}
	return dp[m][n]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
