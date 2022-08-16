package linkedlist

// https://leetcode.cn/problems/merge-in-between-linked-lists/
// 合并两个链表

// 直接a,b索引位置
// 时间复杂度：O(N)
// 空间复杂度：O(1)
func mergeInBetween(list1 *ListNode, a int, b int, list2 *ListNode) *ListNode {
	dummyHead := &ListNode{Next: list1}
	nodea := dummyHead
	for i := 0; i < a; i++ {
		nodea = nodea.Next
	}
	nodeb := nodea
	for i := a; i <= b; i++ {
		nodeb = nodeb.Next
	}
	nodea.Next = list2
	prev := list2
	for node := list2; node != nil; node = node.Next {
		prev = node
	}
	prev.Next = nodeb.Next
	return list1
}
