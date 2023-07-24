package monotonicstack

// https://leetcode.cn/problems/daily-temperatures/
// 每日温度

// 方法一: 单调栈 找第一个比他大的数
// 时间复杂度O(n)
// 空间复杂度O(n)
func dailyTemperatures(temperatures []int) []int {
	ans := make([]int, len(temperatures))
	var stack []int
	for i, temperature := range temperatures {
		ans[i] = 0
		for len(stack) > 0 && temperatures[stack[len(stack)-1]] < temperature {
			ans[stack[len(stack)-1]] = i - stack[len(stack)-1]
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, i)
	}
	return ans
}
