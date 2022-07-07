package main

func longestCommonSubsequence(text1 string, text2 string) int {
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
