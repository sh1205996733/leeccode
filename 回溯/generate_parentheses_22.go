package backtrack

// https://leetcode.cn/problems/generate-parentheses/
// 括号生成

// 方法一：回溯法
// 时间复杂度：O(log)
// 空间复杂度：O(1)
// ["((()))","(()())","(())()","()(())","()()()"]
func generateParenthesis(n int) []string {
	var ans []string
	ans = dfs("", ans, n, n, n*2)
	return ans
}

// 深度优先
func dfs(t string, ans []string, left, right, length int) []string {
	if len(t) == length {
		ans = append(ans, t)
		return ans
	}
	if left > 0 { // 左括弧有剩余 则追加一个左括弧
		ans = dfs(t+"(", ans, left-1, right, length)
	}
	if right > left { // 右括弧有剩余 则追加一个右括弧
		ans = dfs(t+")", ans, left, right-1, length)
	}
	return ans
}
