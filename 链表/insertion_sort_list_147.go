package linkedlist

// https://leetcode.cn/problems/insertion-sort-list/
// 对链表进行插入排序

// 方法一：插入排序 找到插入点tail，如果head>=tail 将tail更新为head head=head.Next
// 如果head比tail小，就从dummyHead开始遍历找到插入点target（第一个大于head的,稳定排序）插入head tail不变 head=head.Next
// 时间复杂度：O(N^2)
// 空间复杂度：O(1)
func insertionSortList(head *ListNode) *ListNode {
	dummyHead := &ListNode{-5001, head}
	tail := dummyHead
	for head != nil {
		if tail.Val <= head.Val {
			tail.Next = head
			tail = head
		} else {
			target := dummyHead
			for ; target.Next.Val <= head.Val; target = target.Next {
			}
			tail.Next = head.Next
			head.Next = target.Next
			target.Next = head
		}
		head = tail.Next
	}
	return dummyHead.Next
}
