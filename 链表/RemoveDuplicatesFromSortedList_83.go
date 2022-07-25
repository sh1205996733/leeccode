package linkedlist

// https://leetcode.cn/problems/remove-duplicates-from-sorted-list/
// 删除排序链表中的重复元素

// 方法一: 迭代
// 时间复杂度为 O(n)
// 空间复杂度：O(1)
func deleteDuplicates0(head *ListNode) *ListNode {
	cur := head
	for cur != nil && cur.Next != nil {
		if cur.Val == cur.Next.Val {
			cur.Next = cur.Next.Next
		} else {
			cur = cur.Next
		}
	}
	return head
}

// 方法二: 递归
// 时间复杂度为 O(n)
// 空间复杂度：O(n)
func deleteDuplicates(head *ListNode) *ListNode {
	if head.Next == nil {
		return head
	}
	head.Next = deleteDuplicates(head.Next)
	if head.Val == head.Next.Val {
		return head.Next
	}
	return head
}
