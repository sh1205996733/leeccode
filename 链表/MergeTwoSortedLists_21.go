package linkedlist

// https://leetcode.cn/problems/merge-two-sorted-lists/
// 合并两个有序链表

// 方法一：栈 双指针，比较俩指针大小，小的指针加入新的新的节点，向前挪 直到有一个指针走完，把另一个指针剩余部分都加到新的节点后面
// 时间复杂度：O(n + m)
// 空间复杂度：O(1)
func mergeTwoLists1(list1 *ListNode, list2 *ListNode) *ListNode {
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

// 方法一：递归 判断 l1 和 l2 哪一个链表的头节点的值更小，然后递归地决定下一个添加到结果里的节点
// 时间复杂度：O(n + m)
// 空间复杂度：O(n + m)
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil {
		return list2
	}
	if list2 == nil {
		return list1
	}
	if list1.Val < list2.Val {
		list1.Next = mergeTwoLists(list1.Next, list2)
		return list1
	} else {
		list2.Next = mergeTwoLists(list1, list2.Next)
		return list2
	}
}
