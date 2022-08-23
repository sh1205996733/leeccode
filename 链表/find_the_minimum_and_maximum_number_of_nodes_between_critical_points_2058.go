package linkedlist

// https://leetcode.cn/problems/find-the-minimum-and-maximum-number-of-nodes-between-critical-points/
// 找出临界点之间的最小和最大距离

// 方法一： 暴力法 滑动窗口
// 最小距离一定出现在两个相邻的临界点之间；
// 最大距离一定出现在第一个和最后一个临界点之间
// 时间复杂度：O(N)
// 空间复杂度：O(1)
func nodesBetweenCriticalPoints(head *ListNode) []int {
	prev, cur, next := head, head.Next, head.Next.Next
	max, min, first, last := -1, -1, -1, -1
	for index := 0; next != nil; index++ {
		if (cur.Val > prev.Val && cur.Val > next.Val) ||
			(cur.Val < prev.Val && cur.Val < next.Val) { //极大值
			if first == -1 {
				first = index
			} else {
				max = index - first
				val := index - last
				if min == -1 || val < min {
					min = val
				}
			}
			last = index
		}
		prev, cur, next = cur, next, next.Next
	}
	return []int{min, max}
}
