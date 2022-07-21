package linkedlist

// https://leetcode.cn/problems/reverse-nodes-in-k-group/
//  K 个一组翻转链表

// 方法一:
// 时间复杂度为 O(n^2)
// 空间复杂度：O(1)
func reverseKGroup(head *ListNode, k int) *ListNode {
	dammy := &ListNode{0, head}
	cur := dammy
	for cur != nil {
		start := cur
		i := k
		for i > 0 && cur != nil {
			cur = cur.Next
			i--
		}
		if i == 0 && cur != nil {
			cur = reverse(start, cur)
		}
	}
	return dammy.Next
}

// 翻转链表 0, 1, 2, 3  --> 0, 3, 2, 1  1
func reverse(head, tail *ListNode) *ListNode {
	cur := head.Next
	newHead := tail.Next
	for cur != tail {
		next := cur.Next
		cur.Next = newHead
		newHead = cur
		cur = next
	}
	cur.Next = newHead
	newHead = head.Next
	head.Next = cur
	return newHead
}
