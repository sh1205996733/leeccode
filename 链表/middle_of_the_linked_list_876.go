package linkedlist

// https://leetcode.cn/problems/middle-of-the-linked-list/
// 链表的中间结点

// 方法一：todo
// 时间复杂度：O(N)
// 空间复杂度：O(1)
func middleNode1(head *ListNode) *ListNode {
	node := head
	size := 0
	for node != nil {
		node = node.Next
		size++
	}
	node = head
	mid := size >> 1
	for i := 1; i <= mid; i++ {
		node = node.Next
	}
	return node
}

// 方法二：快慢指针 slow每次走一步，fast每次走两步，fast走完的时候，slow只走了一半
// 时间复杂度：O(N)
// 空间复杂度：O(1)
func middleNode(head *ListNode) *ListNode {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow
}

// 方法三：将链表转化成数组
// 时间复杂度：O(N)
// 空间复杂度：O(N)
