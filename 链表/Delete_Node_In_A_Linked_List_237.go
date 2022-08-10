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
