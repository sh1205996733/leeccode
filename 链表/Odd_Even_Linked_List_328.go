package linkedlist

// https://leetcode.cn/problems/odd-even-linked-list/
// 奇偶链表

// 方法一：迭代 一次跳两步，再加一个偶数头(有点像DNA螺旋)
// 时间复杂度：O(1)
// 空间复杂度：O(N)
func oddEvenList(head *ListNode) *ListNode {
	//if head == nil {
	//	return head
	//}
	//head2 := head.Next
	//node, node2 := head, head2
	//for node != nil && node.Next != nil {
	//	next := node.Next
	//	node.Next = next.Next
	//	next.Next = nil
	//	node2.Next = next
	//	node2 = next
	//	last = node
	//	node = node.Next
	//}
	//if node != nil {
	//	last = node
	//}
	//last.Next = head2
	//return head

	if head == nil {
		return head
	}
	evenHead := head.Next
	odd := head
	even := evenHead
	for even != nil && even.Next != nil {
		odd.Next = even.Next
		odd = odd.Next
		even.Next = odd.Next
		even = even.Next
	}
	odd.Next = evenHead
	return head
}
