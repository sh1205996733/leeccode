package binarytree

import "fmt"

// TreeNode Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func (p *TreeNode) RootNode() interface{} {
	return p
}
func (p *TreeNode) LNode() interface{} {
	return p.Left
}
func (p *TreeNode) RNode() interface{} {
	return p.Right
}
func (p *TreeNode) ToString() interface{} {
	return p.Val
}
func (p *TreeNode) ColorOf() bool {
	return true
}

// MakeLinkedList2 生成链表
func MakeLinkedList2(nums []int) *TreeNode {
	var head, tail *TreeNode
	for _, num := range nums {
		node := &TreeNode{num, nil, nil}
		if head == nil {
			head = node
		} else {
			tail.Right = node
		}
		tail = node
	}
	PrintLinkedList2(head)
	return head
}

// PrintLinkedList2 打印链表
func PrintLinkedList2(node *TreeNode) {
	str := ""
	for node != nil {
		str += fmt.Sprintf("%d->", node.Val)
		node = node.Right
	}
	if len(str) > 0 {
		str = str[:len(str)-2]
	}
	fmt.Println(str)
}
