package linkedlist

// https://leetcode.cn/problems/delete-node-in-a-linked-list/
// 删除链表中的节点

// 方法一：和下一个节点交换 直接删除 保证需要删除的节点 不是末尾节点
// 时间复杂度：O(1)
// 空间复杂度：O(1)
func deleteNode(node *ListNode) {
	node.Val = node.Next.Val
	node.Next = node.Next.Next
}

// https://leetcode.cn/problems/shan-chu-lian-biao-de-jie-dian-lcof/
func deleteNode1(head *ListNode, val int) *ListNode {
	if head == nil {
		return nil
	}
	if head.Val == val {
		return head.Next
	}
	prev, cur := head, head.Next
	for cur != nil && cur.Val != val {
		prev = cur
		cur = cur.Next
	}
	if cur != nil {
		prev.Next = cur.Next
	}
	return head
}
