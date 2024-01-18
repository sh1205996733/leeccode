package binarytree

import (
	"container/list"
	"math"
)

// 恢复二叉搜索树
// https://leetcode.cn/problems/recover-binary-search-tree/description/

func recoverTree(root *TreeNode) {
	stack := list.New()
	var first, second *TreeNode
	node := root
	min := &TreeNode{Val: math.MinInt64, Left: nil, Right: nil}
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
