package linkedlist

// https://leetcode.cn/problems/print-immutable-linked-list-in-reverse/
// 逆序打印不可逆链表

// 方法一：栈，先入栈，在出栈
// 时间复杂度：O(N)
// 空间复杂度：O(1)
func printLinkedListInReserve0(head *ListNode) {
	var stack []int
	for head != nil {
		stack = append(stack, head.Val)
		head = head.Next
	}
	for i := len(stack) - 1; i >= 0; i-- {
		print(stack[i])
	}
}

// 方法二：递归
// 时间复杂度：O(N)
// 空间复杂度：O(N)
func printLinkedListInReserve(head *ListNode) {
	if head == nil {
		return
	}
	printLinkedListInReserve(head.Next)
	print(head.Val)
}

// https://leetcode.cn/problems/cong-wei-dao-tou-da-yin-lian-biao-lcof/
// 剑指 Offer 06. 从尾到头打印链表
func reversePrint(head *ListNode) []int {
	var stack []int
	for head != nil {
		stack = append(stack, head.Val)
		head = head.Next
	}
	l := len(stack)
	for i := 0; i < l>>1; i++ {
		stack[i], stack[l-1-i] = stack[l-1-i], stack[i]
	}
	return stack
}
