package main

import "fmt"

func main() {
	head := &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, &ListNode{6, nil}}}}}}
	head = middleNode(head)
	head.string()
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

//常规
func middleNode1(head *ListNode) *ListNode {
	node := head
	size := 0
	for node != nil {
		node = node.Next
		size++
	}
	node = head
	mid := size >> 1
	for i := 1; i <= mid; i++ {
		node = node.Next
	}
	return node
}

//快慢指针 slow每次走一步，fast每次走两步，fast走完的时候，slow只走了一半
func middleNode(head *ListNode) *ListNode {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow
}
