package linkedlist

// https://leetcode.cn/problems/remove-duplicates-from-sorted-list-ii/
// 删除排序链表中的重复元素 II

// 方法一: 迭代 双指针
// 时间复杂度为 O(n)
// 空间复杂度：O(1)
func deleteDuplicatesII(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	pre := head
	cur := head.Next
	var tail *ListNode
	if pre.Val != cur.Val {
		tail = pre
	}
	dummyHead := &ListNode{0, tail}
	for cur != nil && cur.Next != nil {
		if pre.Val != cur.Val && cur.Val != cur.Next.Val {
			if tail == nil {
				dummyHead.Next = cur
			} else {
				tail.Next = cur
			}
			tail = cur
		}
		pre = cur
		cur = cur.Next
	}
	if pre.Val == cur.Val {
		cur = cur.Next
	}
	if tail != nil {
		tail.Next = cur
	}
	return dummyHead.Next
}

// 方法二: 递归
// 时间复杂度为 O(n)
// 空间复杂度：O(n)
func deleteDuplicatesII1(head *ListNode) *ListNode {
	if head.Next == nil {
		return head
	}
	head.Next = deleteDuplicates(head.Next)
	if head.Val == head.Next.Val {
		return head.Next.Next
	}
	return head
}
