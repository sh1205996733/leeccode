package linkedlist

// https://leetcode.cn/problems/remove-duplicates-from-sorted-list-ii/
// 删除排序链表中的重复元素 II

// 方法一: 迭代 双指针 找到相同节点的前驱prev和后继节点next
// 时间复杂度为 O(n)
// 空间复杂度：O(1)
func deleteDuplicatesII(head *ListNode) *ListNode {
	dummyHead := &ListNode{101, head}
	prev := dummyHead
	cur := head
	for cur != nil {
		next := cur.Next
		for ; next != nil && cur.Val == next.Val; next = next.Next {
		}
		if cur.Next == next {
			prev = cur
		} else {
			prev.Next = next
		}
		cur = next
	}
	return dummyHead.Next
}

// 方法二: 递归
// 时间复杂度为 O(n)
// 空间复杂度：O(n)
func deleteDuplicatesII0(head *ListNode) *ListNode {
	if head.Next == nil {
		return head
	}
	tail := deleteDuplicatesII0(head.Next)
	if head.Val == head.Next.Val { //前后俩个节点值相同
		if head.Next == tail { //tail就是head的next
			tail = tail.Next
		}
		return tail
	} else {
		head.Next = tail
	}
	return head
}
