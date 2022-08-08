package linkedlist

// https://leetcode.cn/problems/reverse-linked-list-ii/
// 反转链表

// 方法一: 迭代 先找到left-1、right的位置，再使用头插法交换位置
// 时间复杂度为 O(n)
// 空间复杂度：O(1)
func reverseBetween(head *ListNode, left int, right int) *ListNode {
	dammyHead := &ListNode{0, head}
	cur := dammyHead
	prev := cur
	for i := 0; i < right; i++ {
		if i < left {
			prev = cur
			cur = cur.Next
			continue
		}
		next := cur.Next
		cur.Next = next.Next
		next.Next = prev.Next
		prev.Next = next
	}
	return dammyHead.Next
}
