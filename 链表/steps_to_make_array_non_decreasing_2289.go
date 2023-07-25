package linkedlist

// https://leetcode.cn/problems/steps-to-make-array-non-decreasing/
// 使数组按非递减顺序排列
// 5,3,4,4,7,3,6,11,8,5,11
// 5,4,4,7,6,11,11
// 5,4,7,11,11
// 5,7,11,11

// [5,3,4,4,2,3,6]
// [6]
// [5,6]
// 方法一：单调栈
// 时间复杂度：O(N)
// 空间复杂度：O(n)
func totalSteps(nums []int) int {
	type pair struct{ v, t int }
	st, ans := []pair{}, 0
	for _, num := range nums {
		maxT := 0
		for len(st) > 0 && st[len(st)-1].v <= num {
			maxT = max(maxT, st[len(st)-1].t)
			st = st[:len(st)-1]
		}
		if len(st) > 0 {
			maxT++
			ans = max(ans, maxT)
		} else {
			maxT = 0
		}
		st = append(st, pair{num, maxT})
	}
	return ans
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
