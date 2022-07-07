package main

import "fmt"

func main() {
	head := &ListNode{1, &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{3, &ListNode{4, nil}}}}}}
	head = deleteDuplicates(head)
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

//迭代
func deleteDuplicates1(head *ListNode) *ListNode {
	node := head
	for node != nil && node.Next != nil {
		if node.Val == node.Next.Val { //删除下一个
			node.Next = node.Next.Next
		} else {
			node = node.Next
		}
	}
	return head
}

//递归
func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	head.Next = deleteDuplicates(head.Next)
	if head.Val == head.Next.Val {
		return head.Next
	}
	return head
}
