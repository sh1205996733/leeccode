package linkedlist

// https://leetcode.cn/problems/add-two-numbers-ii/
// https://leetcode.cn/problems/lMSNwu/submissions/
// 两数相加 II

// 方法一：先反转，再计算,最后再将结果反转
// 时间复杂度：O(N)
// 空间复杂度：O(1)
func addTwoNumbersII0(l1 *ListNode, l2 *ListNode) *ListNode {
	l1 = reserve(l1)
	l2 = reserve(l2)
	carry := 0
	var head, tail *ListNode
	sum := 0
	// 如果l1、l2都不为空,就进行运算 注意加上进位carry
	// 如果l1不为空 说明l2为空 则将l1剩余部分继续做运算，注意加上进位carry
	// 如果l2不为空 说明l1为空 则将l2剩余部分继续做运算，注意加上进位carry
	for l1 != nil || l2 != nil {
		var val1, val2 int
		if l1 != nil {
			val1 = l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			val2 = l2.Val
			l2 = l2.Next
		}
		sum = val1 + val2 + carry
		carry = sum / 10
		node := &ListNode{sum % 10, nil}
		if head == nil {
			head = node
		} else {
			tail.Next = node
		}
		tail = node
	}
	// 最后还要判断一下进位carry
	if carry > 0 {
		tail.Next = &ListNode{carry, nil}
	}
	return reserve(head)
}

func reserve(head *ListNode) *ListNode {
	dummyHead := &ListNode{}
	for head != nil {
		next := head.Next
		head.Next = dummyHead.Next
		dummyHead.Next = head
		head = next
	}
	return dummyHead.Next
}

// 方法二：栈 三个栈
// 时间复杂度：O(N)
// 空间复杂度：O(N)
func addTwoNumbersII(l1 *ListNode, l2 *ListNode) *ListNode {
	dummyHead := &ListNode{}
	var stack1, stack2 []int
	for ; l1 != nil; l1 = l1.Next {
		stack1 = append(stack1, l1.Val)
	}
	for ; l2 != nil; l2 = l2.Next {
		stack2 = append(stack2, l2.Val)
	}
	sum, carry := 0, 0
	for len(stack1) > 0 || len(stack2) > 0 {
		var val1, val2 int
		if len(stack1) > 0 {
			val1 = stack1[len(stack1)-1]
			stack1 = stack1[:len(stack1)-1]
		}
		if len(stack2) > 0 {
			val2 = stack2[len(stack2)-1]
			stack2 = stack2[:len(stack2)-1]
		}
		sum = val1 + val2 + carry
		carry = sum / 10
		dummyHead.Next = &ListNode{sum % 10, dummyHead.Next}
	}
	// 最后还要判断一下进位carry
	if carry > 0 {
		dummyHead.Next = &ListNode{carry, dummyHead.Next}
	}
	return dummyHead.Next
}
