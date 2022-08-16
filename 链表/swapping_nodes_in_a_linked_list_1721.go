package linkedlist

// https://leetcode.cn/problems/swapping-nodes-in-a-linked-list/
// 交换链表中的节点

// 双指针
// 时间复杂度：O(N)
// 空间复杂度：O(1)
func swapNodes(head *ListNode, k int) *ListNode {
	if head == nil || head.Next == nil || k == 0 {
		return head
	}
	fast := head
	slow := head
	var first *ListNode
	// 找到正数和倒数第 k 个结点
	i := 1
	for fast != nil {
		if i <= k {
			first = fast
			i++
		} else {
			slow = slow.Next
		}
		fast = fast.Next
	}
	// 交换数据域
	first.Val, slow.Val = slow.Val, first.Val
	return head
}
