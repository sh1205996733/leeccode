package main

import (
	"container/list"
	"math"
)

func main() {
	minDiffInBST(nil)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//非递归-栈实现
func minDiffInBST(root *TreeNode) int {
	min := math.MaxInt64
	var minVal = -1
	stack := list.New()
	node := root
	for node != nil || stack.Len() > 0 {
		if node != nil {
			stack.PushBack(node)
			node = node.Left
		} else {
			node = stack.Remove(stack.Back()).(*TreeNode)
			abs := node.Val - minVal
			if minVal != -1 {
				if abs < min {
					min = abs
				}
			}
			minVal = node.Val
			node = node.Right
		}
	}
	return min
}
