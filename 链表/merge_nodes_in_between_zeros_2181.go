package linkedlist

// https://leetcode.cn/problems/merge-nodes-in-between-zeros/
// 合并零之间的节点

func mergeNodes(head *ListNode) *ListNode {
	dummyHead := &ListNode{}
	tail := dummyHead
	t := head
	for node := head.Next; node != nil; node = node.Next {
		if node.Val != 0 {
			t.Val += node.Val
		} else {
			t.Next = nil
			tail.Next = t
			tail = tail.Next
			t = node
		}
	}
	return dummyHead.Next
}
