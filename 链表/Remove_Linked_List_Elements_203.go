package linkedlist

// https://leetcode.cn/problems/remove-linked-list-elements/
// 移除链表元素

// 方法一：迭代+虚拟头节点
// 时间复杂度：O(N)
// 空间复杂度：O(1)
func removeElements0(head *ListNode, val int) *ListNode {
	dammyHead := &ListNode{-1, head}
	for node := dammyHead; node.Next != nil; {
		if node.Next.Val == val {
			node.Next = node.Next.Next
		} else {
			node = node.Next
		}
	}
	return dammyHead.Next
}

// 方法二：递归
// 时间复杂度：O(N)
// 空间复杂度：O(N)
func removeElements(head *ListNode, val int) *ListNode {
	if head == nil {
		return nil
	}
	head.Next = removeElements(head.Next, val)
	if head.Val == val {
		return head.Next
	}
	return head
}
