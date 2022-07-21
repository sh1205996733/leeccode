package linkedlist

// https://leetcode.cn/problems/swap-nodes-in-pairs/
// 两两交换链表中的节点

//方法一: 迭代 两两交换值 步长为2 直到下一个为空或者下一个的Next为空结束
// 方法一：两两合并 合并两个有序链表(不交换节点)
// 时间复杂度为 O(n)
// 空间复杂度：O(1)
func swapPairs0(head *ListNode) *ListNode {
	cur := head
	for cur != nil && cur.Next != nil {
		cur.Val, cur.Next.Val = cur.Next.Val, cur.Val
		cur = cur.Next.Next
	}
	return head
}

// 迭代 只能交换节点 1-2-3-4-5
func swapPairs1(head *ListNode) *ListNode {
	if head == nil || head.Next == nil { //只有一个元素就直接返回
		return head
	}
	cur := head
	prev := cur
	head = head.Next
	for cur != nil && cur.Next != nil {
		next := cur.Next
		prev.Next = next
		cur.Next = next.Next
		next.Next = cur
		prev = cur
		cur = cur.Next
	}
	return head
}

// 迭代 只能交换节点 0-2-1-3-4-5 (虚拟头结点)
func swapPairs2(head *ListNode) *ListNode {
	dummy := &ListNode{0, head}
	cur := dummy.Next
	prev := dummy
	for cur != nil && cur.Next != nil {
		next := cur.Next
		prev.Next = next
		cur.Next = next.Next
		next.Next = cur
		prev = cur
		cur = cur.Next
	}
	return dummy.Next
}

//方法二: 递归
// 时间复杂度为 O(n)
// 空间复杂度：O(1)
func swapPairs3(head *ListNode) *ListNode {
	if head == nil || head.Next == nil { //只有一个元素就直接返回
		return head
	}
	swapPairs(head.Next.Next)
	head.Val, head.Next.Val = head.Next.Val, head.Val
	return head
}

// 递归 只能交换节点 1-2-4-3-5
func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil { //只有一个元素就直接返回
		return head
	}
	next := head.Next
	head.Next = swapPairs(next.Next)
	next.Next = head
	return next
}
