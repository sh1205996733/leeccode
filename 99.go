package main

import (
	"container/list"
	"math"
)

func main() {
	recoverTree(nil)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func recoverTree(root *TreeNode) {
	stack := list.New()
	var first, second *TreeNode
	node := root
	min := &TreeNode{math.MinInt64, nil, nil}
	for node != nil || stack.Len() > 0 {
		if node != nil {
			stack.PushBack(node)
			node = node.Left
		} else { //左边为空
			node = stack.Remove(stack.Back()).(*TreeNode)
			// 访问节点
			if first == nil && min.Val > node.Val {
				first = min
			}
			if first != nil && min.Val > node.Val {
				second = node
			}
			min = node
			node = node.Right
		}
	}
	first.Val, second.Val = second.Val, first.Val
}
