package linkedlist

// https://leetcode.cn/problems/reverse-linked-list/
// 反转链表

// 方法一：迭代 头插法
// 时间复杂度：O(N)
// 空间复杂度：O(1)
func reverseList0(head *ListNode) *ListNode {
	dammyHead := &ListNode{}
	for head != nil {
		next := head.Next
		head.Next = dammyHead.Next
		dammyHead.Next = head
		head = next
	}
	return dammyHead.Next
}

// 方法二：递归
// 时间复杂度：O(N)
// 空间复杂度：O(N)
func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	newHead := reverseList(head.Next)
	head.Next.Next = head
	head.Next = nil
	return newHead
}
