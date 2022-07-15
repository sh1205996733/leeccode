package linkedlist

// 删除链表的倒数第 N 个结点
// https://leetcode.cn/problems/remove-nth-node-from-end-of-list

// 方法一: 获取链表的长度,再找到正数N+1个节点target，然后将target的Next指向target.Next.Next 注意最后一个节点
//时间复杂度：O(L)O(L)，其中 LL 是链表的长度。
//空间复杂度：O(1)O(1)。
func removeNthFromEnd0(head *ListNode, n int) *ListNode {
	cur := head
	size := 0
	// 获取链表的长度
	for ; cur != nil; cur = cur.Next {
		size++
	}
	dummy := &ListNode{0, head}
	cur = dummy
	// 找到第N+1的节点
	for i := 0; i < size-n; i++ {
		cur = cur.Next
	}
	cur.Next = cur.Next.Next
	return dummy.Next
}

// 方法二：栈 先将head依次入栈，找到下标N的前驱节点
//时间复杂度：O(L)O(L)，其中 LL 是链表的长度。
//空间复杂度：O(L)O(L)，其中 LL 是链表的长度。主要为栈的开销。
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	stack := []*ListNode{}
	dummy := &ListNode{0, head}
	for node := dummy; node != nil; node = node.Next {
		stack = append(stack, node)
	}
	prev := stack[len(stack)-n-1]
	prev.Next = prev.Next.Next
	return dummy.Next
}

// 方法三：快慢指针 快慢指针从间隔N个节点开始往下走 直到快指针结束，慢指针的位置就是Target
//时间复杂度：O(L)O(L)，其中 LL 是链表的长度。
//空间复杂度：O(1)O(1)。
func removeNthFromEnd1(head *ListNode, n int) *ListNode {
	dummy := &ListNode{0, head}
	fast := head
	for i := 0; i < n; i++ {
		fast = fast.Next
	}
	slow := dummy
	for fast != nil {
		fast = fast.Next
		slow = slow.Next
	}
	slow.Next = slow.Next.Next
	return dummy.Next
}
