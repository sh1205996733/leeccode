package linkedlist

import "fmt"

// ListNode Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

// MakeLinkedList 生成链表
func MakeLinkedList(nums []int) *ListNode {
	var head, tail *ListNode
	for _, num := range nums {
		node := &ListNode{num, nil}
		if head == nil {
			head = node
		} else {
			tail.Next = node
		}
		tail = node
	}
	PrintLinkedList(head)
	return head
}

// ParseLinkedList int转成链表
func ParseLinkedList(n int) *ListNode {
	var head, tail *ListNode
	for n > 0 {
		node := &ListNode{n % 10, nil}
		if head == nil {
			head = node
		} else {
			tail.Next = node
		}
		tail = node
		n = n / 10
	}
	return head
}

// ParseInt 链表转数字(仅int范围)
func ParseInt(node *ListNode) int {
	cur := node
	result, base := 0, 1
	for cur != nil {
		result += cur.Val * base
		base *= 10
		cur = cur.Next
	}
	return result
}

// PrintLinkedList 打印链表
func PrintLinkedList(node *ListNode) {
	str := ""
	for node != nil {
		str += fmt.Sprintf("%d->", node.Val)
		node = node.Next
	}
	if len(str) > 0 {
		str = str[:len(str)-2]
	}
	fmt.Println(str)
}

func GetMidNode(head *ListNode) interface{} {
	if head == nil || head.Next == nil {
		return head
	}
	slow := head
	fast := head.Next.Next
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}
	//return slow //偶数取size/2 -1
	return slow.Next //偶数取size/2
}
