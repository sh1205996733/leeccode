package linkedlist

// https://leetcode.cn/problems/sort-list/
// 方法一：先找出中间元素，把链表分成两段，再递归对每一段进行前两步操作，直到不能再分，进行把每段有序链表合并（分治思想、归并排序）
func sortList(head *ListNode) *ListNode {
	return sortSubList(head, nil)
}

func sortSubList(head *ListNode, tail *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	mid := getMidNode(head, tail)
	return mergeTwoLists2(sortSubList(head, mid), sortSubList(mid, tail))
}

func mergeTwoLists2(list1 *ListNode, list2 *ListNode) *ListNode {
	first := list1
	second := list2
	dummy := &ListNode{0, nil}
	tail := dummy
	for first != nil && second != nil {
		if first.Val > second.Val {
			tail.Next = second
			second = second.Next
		} else {
			tail.Next = first
			first = first.Next
		}
		tail = tail.Next
	}
	if first != nil {
		tail.Next = first
	}
	if second != nil {
		tail.Next = second
	}
	return dummy.Next
}

func getMidNode(head, tail *ListNode) *ListNode {
	if head == tail || head.Next == tail {
		return head
	}
	slow := head
	fast := head.Next.Next
	for fast != tail && fast.Next != tail {
		fast = fast.Next.Next
		slow = slow.Next
	}
	return slow.Next
}
