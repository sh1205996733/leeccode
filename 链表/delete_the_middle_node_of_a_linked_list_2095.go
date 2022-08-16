package linkedlist

// https://leetcode.cn/problems/delete-the-middle-node-of-a-linked-list/
// 删除链表的中间节点

// 双指针
// 时间复杂度：O(N)
// 空间复杂度：O(1)
func deleteMiddle(head *ListNode) *ListNode {
	dummyHead := &ListNode{Next: head}
	prev, slow, fast := dummyHead, head, head
	for fast != nil && fast.Next != nil {
		prev = slow
		slow = slow.Next
		fast = fast.Next.Next
	}
	prev.Next = slow.Next
	return dummyHead.Next
}
