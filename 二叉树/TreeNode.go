package binarytree

import (
	"fmt"
	"leetcode/二叉树/printer"
)

// TreeNode Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
	Next  *TreeNode
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
//func MakeLinkedList2(nums []int) *TreeNode {
//	var head, tail *TreeNode
//	for _, num := range nums {
//		node := &TreeNode{num, nil, nil}
//		if head == nil {
//			head = node
//		} else {
//			tail.Right = node
//		}
//		tail = node
//	}
//	PrintLinkedList2(head)
//	return head
//}

// PrintLinkedList2 打印链表
func PrintLinkedList2(node *TreeNode) {
	str := ""
	var head *TreeNode
	for node != nil {
		if head == nil {
			head = node.Left
		}
		str += fmt.Sprintf("%d->", node.Val)
		node = node.Next
		if node == nil {
			node = head
			head = nil
		}
	}
	if len(str) > 0 {
		str = str[:len(str)-2]
	}
	fmt.Println(str)
}

// MakePerfectBinaryTree 完美二叉树
func MakePerfectBinaryTree(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	var root *TreeNode
	m := map[int]*TreeNode{}
	for i, num := range nums {
		m[i] = &TreeNode{Val: num}
	}
	for i, _ := range nums {
		if i == 0 {
			root = m[i]
		}
		m[i].Left = m[2*i+1]
		m[i].Right = m[2*i+2]
	}
	printer.Println(root)
	return root
}
