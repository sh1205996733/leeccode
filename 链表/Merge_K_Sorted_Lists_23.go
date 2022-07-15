package linkedlist

// https://leetcode.cn/problems/merge-k-sorted-lists/
// 合并K个升序链表

// 方法一：两两合并 合并两个有序链表
// 时间复杂度为 O(k^2 n)
// 空间复杂度：O(1)
func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}
	head := lists[0]
	for i := 1; i < len(lists); i++ {
		head = merge2Lists(head, lists[i])
	}
	return head
}

// 方法二：使用优先队列合并(最小堆)
// 时间复杂度：考虑优先队列中的元素不超过 kk 个，那么插入和删除的时间代价为 O(logk)，这里最多有 knkn 个点，对于每个点都被插入删除各一次，故总的时间代价即渐进时间复杂度为 O(kn×logk)。
// 空间复杂度：这里用了优先队列，优先队列中的元素不超过 kk 个，故渐进空间复杂度为 O(k)。
func mergeKLists2(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}
	head := lists[0]
	for i := 1; i < len(lists); i++ {
		head = merge2Lists(head, lists[i])
	}
	return head
}
func merge2Lists(list1 *ListNode, list2 *ListNode) *ListNode {
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
