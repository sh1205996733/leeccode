package main

import (
	"fmt"
	"leetcode/print/printer"
)

func main() {
	root := &TreeNode{5,
		&TreeNode{3,
			&TreeNode{2, nil, nil},
			&TreeNode{4, nil, nil}},
		&TreeNode{7,
			&TreeNode{6, nil, nil},
			&TreeNode{8, nil, nil}}}
	printer.Println(root)
	deleteNode(root, 3)
	fmt.Println("---------删除后---------")
	printer.Println(root)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func (p *TreeNode) Root() interface{} {
	return p
}
func (p *TreeNode) LNode() interface{} {
	if p.Left == nil {
		return nil
	}
	return p.Left
}
func (p *TreeNode) RNode() interface{} {
	if p.Right == nil {
		return nil
	}
	return p.Right
}
func (p *TreeNode) ToString() interface{} {
	return p.Val
}
func (p *TreeNode) ColorOf() bool {
	return false
}
func successor(root *TreeNode) int {
	root = root.Right
	for root.Left != nil {
		root = root.Left
	}
	return root.Val
}
func predecessor(root *TreeNode) int {
	root = root.Left
	for root.Right != nil {
		root = root.Right
	}
	return root.Val
}

func deleteNode(root *TreeNode, key int) *TreeNode {
	if root == nil {
		return root
	}
	if key > root.Val {
		root.Right = deleteNode(root.Right, key)
	} else if key < root.Val {
		root.Left = deleteNode(root.Left, key)
	} else {
		if root.Left == nil && root.Right == nil {
			return nil
		} else if root.Right != nil {
			root.Val = successor(root)
			root.Right = deleteNode(root.Right, root.Val)
		} else {
			root.Val = predecessor(root)
			root.Left = deleteNode(root.Left, root.Val)
		}
	}
	return root
}
