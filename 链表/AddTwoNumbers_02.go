package linkedlist

import "fmt"

// https://leetcode.cn/problems/add-two-numbers/
// eg:l1 = [2,4,3], l2 = [5,6,4] [7,0,8] 342 + 465 = 807

//方法一: 暴力法 将l1和l2转化成整数,在将结果转成链表,再相加但结果会溢出
// 时间复杂度O(3n) 空间复杂度O(n)
func addTwoNumbers1(l1 *ListNode, l2 *ListNode) *ListNode {
	//l1和l2都为nil时,直接返回
	num1 := ParseInt(l1)
	num2 := ParseInt(l2)
	fmt.Println("num1 = ", num1, "num2 = ", num2)
	return ParseLinkedList(num1 + num2)
}

//方法二: 就地运算法：双指针一个遍历l1一个遍历l2
// 时间复杂度O(n) 空间复杂度O(n)
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
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
	return head
}
