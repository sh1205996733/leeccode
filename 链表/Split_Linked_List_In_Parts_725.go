package linkedlist

import "math"

// https://leetcode.cn/problems/split-linked-list-in-parts/
// 分隔链表

// 方法一：迭代 先计算长度count，然后count/k 向上取整，正数即为数量，直到k==0
// 时间复杂度：O(1)
// 空间复杂度：O(N)
func splitListToParts(head *ListNode, k int) []*ListNode {
	node := head
	count := 0
	res := make([]*ListNode, k)
	for node != nil {
		count++
		node = node.Next
	}
	for ; k > 0 && count > 0; k-- {
		n := int(math.Ceil(float64(count) / float64(k)))
		node = head
		for i := n; i > 1; i-- {
			node = node.Next
		}
		res[cap(res)-k] = head
		head, node.Next = node.Next, nil
		count -= n
	}
	return res
}
