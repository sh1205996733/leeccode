package main

import "fmt"

func main() {
	e1 := &ListNode{4, nil}
	e2 := &ListNode{5, e1}
	e1.Next = e2
	head := &ListNode{1, &ListNode{2, &ListNode{3, e1}}}
	fmt.Println(hasCycle(head))
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func (node *ListNode) string() {
	for node != nil {
		fmt.Printf("%d_", node.Val)
		node = node.Next
	}
	fmt.Println()
}

func hasCycle(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}
	slow, fast := head, head.Next //不能一开始就相遇,所以slow,fast不能一开始就相等
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if fast == slow {
			return true
		}
	}
	return false
}
