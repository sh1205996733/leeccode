package main

import "fmt"

func main() {
	head := &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, nil}}}}}
	deleteNode(head)
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
}
func deleteNode(node *ListNode) {
	node.Val = node.Next.Val
	node.Next = node.Next.Next
}
