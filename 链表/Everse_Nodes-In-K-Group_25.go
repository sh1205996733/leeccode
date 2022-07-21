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
			reverse(start, cur.Next)
		}
	}
	return dammy.Next
}

// 翻转链表 0-1-2-4-5-6-7-8
func reverse(head, end *ListNode) {
	cur := head.Next
	var newHead *ListNode
	for cur != end {
		next := cur.Next
		cur.Next = newHead
		newHead = cur
		cur = next
	}
	head.Next = newHead
}
