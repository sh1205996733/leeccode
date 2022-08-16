package linkedlist

// https://leetcode.cn/problems/sort-list/
// https://leetcode.cn/problems/7WHec2/
// 排序链表
// 方法一：先找出中间元素，把链表分成两段，再递归对每一段进行前两步操作，直到不能再分，进行把每段有序链表合并（分治思想、归并排序）
//时间复杂度：O(nlogn)，其中 nn 是链表的长度。
//空间复杂度：O(logn)，其中 nn 是链表的长度。空间复杂度主要取决于递归调用的栈空间。
func sortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	mid := getMidNode(head)
	tail := mid.Next
	mid.Next = nil
	return mergeTwoLists2(sortList(head), sortList(tail))
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

func getMidNode(head *ListNode) *ListNode {
	slow := head
	fast := head.Next.Next
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}
	return slow
}
