package linkedlist

// https://leetcode.cn/problems/remove-duplicates-from-sorted-list-ii/
// 删除排序链表中的重复元素 II

// 方法一: 迭代 双指针
// 时间复杂度为 O(n)
// 空间复杂度：O(1)
func deleteDuplicatesII0(head *ListNode) *ListNode {
	dammyHead := &ListNode{0, nil}
	cur := head
	//tail := head
	for cur != nil && cur.Next != nil {
		if cur.Val != cur.Next.Val {
			dammyHead.Next = cur
		} else {
		}
		cur = cur.Next
	}
	return dammyHead.Next
}

// 方法二: 递归
// 时间复杂度为 O(n)
// 空间复杂度：O(n)
func deleteDuplicatesII(head *ListNode) *ListNode {
	if head.Next == nil {
		return head
	}
	head.Next = deleteDuplicates(head.Next)
	if head.Val == head.Next.Val {
		return head.Next.Next
	}
	return head
}
