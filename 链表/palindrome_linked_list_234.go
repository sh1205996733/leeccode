package linkedlist

// https://leetcode.cn/problems/palindrome-linked-list/
// https://leetcode.cn/problems/palindrome-linked-list-lcci/submissions/
// https://leetcode.cn/problems/aMhZSa/
// 回文链表

// 方法一：迭代 找中间节点mid，从mid处还是反转 然后比较
// 时间复杂度：O(N)
// 空间复杂度：O(1)
func isPalindrome(head *ListNode) bool {
	if head.Next == nil {
		return true
	}
	mid, fast := head, head.Next
	for fast != nil && fast.Next != nil {
		mid = mid.Next
		fast = fast.Next.Next
	}
	dummyHead := &ListNode{}
	node := mid.Next
	mid.Next = nil
	for node != nil {
		next := node.Next
		node.Next = dummyHead.Next
		dummyHead.Next = node
		node = next
	}
	newHead := dummyHead.Next
	for newHead != nil {
		if head.Val != newHead.Val {
			return false
		}
		head = head.Next
		newHead = newHead.Next
	}
	return true
}

// 方法一：转换为回文串，类似方法还有用栈或数组判断
// 时间复杂度：O(N)
// 空间复杂度：O(N)
