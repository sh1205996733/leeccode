package linkedlist

// https://leetcode.cn/problems/reorder-list/
// https://leetcode.cn/problems/LGjMqU/submissions/
// 重排链表

// 先找到中间节点mid 从中间节点mid开始反转获取新的newHead，然后将newHead节点插入到head知道mid时接结束，返回原head
// 时间复杂度：O(N)
// 空间复杂度：O(1)
func reorderList(head *ListNode) {
	slow, fast := head, head.Next
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	var newHead *ListNode
	node := slow.Next
	slow.Next = nil
	for node != nil {
		next := node.Next
		node.Next = newHead
		newHead = node
		node = next
	}
	node = head
	for newHead != nil {
		next := newHead.Next
		newHead.Next = node.Next
		node.Next = newHead
		newHead = next
		node = node.Next.Next
	}
}
