package linkedlist

// https://leetcode.cn/tag/linked-list/problemset/
// 单链表加1
// 用一个 非空 单链表来表示一个非负整数，然后将这个整数加一。你可以假设这个整数除了 0 本身，没有任何前导的 0。
// 这个整数的各个数位按照 高位在链表头部、低位在链表尾部 的顺序排列。

// 方法一：遍历 先反转 在遍历
// 时间复杂度：O(N)
// 空间复杂度：O(1)
func plusOne0(head *ListNode) *ListNode {
	dummyHead := &ListNode{}
	for node := head; node != nil; {
		next := node.Next
		node.Next = dummyHead.Next
		dummyHead.Next = node
		node = next
	}
	// 进位
	carry := 1
	for node := dummyHead.Next; node != nil; node = node.Next {
		carry = (node.Val + carry) / 10
		node.Val = (node.Val + carry) % 10
	}

	dummyHead1 := &ListNode{}
	for node := dummyHead.Next; node != nil; {
		next := node.Next
		node.Next = dummyHead1.Next
		dummyHead1.Next = node
		node = next
	}
	if carry == 1 {
		dummyHead1.Next = &ListNode{1, dummyHead1.Next}
	}
	return dummyHead1.Next
}

// 方法二：递归
// 时间复杂度：O(N)
// 空间复杂度：O(N)
var carry = 1

func plusOne(head *ListNode) *ListNode {
	plusOne2(head)
	if carry == 1 {
		head = &ListNode{1, head}
	}
	return head
}

func plusOne2(head *ListNode) {
	if head == nil {
		return
	}
	plusOne2(head.Next)
	carry = (head.Val + carry) / 10
	head.Val = (head.Val + carry) % 10
}
