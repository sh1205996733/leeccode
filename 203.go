package main

import "fmt"

func main() {
	head := &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, &ListNode{6, nil}}}}}}
	head = removeElements(head, 5)
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

//1. 删除头结点时另做考虑（由于头结点没有前一个结点）
//2. 添加一个虚拟头结点，删除头结点就不用另做考虑
//3. 递归

//删除头结点时另做考虑
func removeElements(head *ListNode, val int) *ListNode {

	for head != nil {
		if head.Val == val {
			head = head.Next
		} else {
			break
		}
	}
	node := head
	for node != nil && node.Next != nil {
		if node.Next.Val == val { //删除下一个
			node.Next = node.Next.Next
		} else {
			node = node.Next
		}
	}
	return head
}
