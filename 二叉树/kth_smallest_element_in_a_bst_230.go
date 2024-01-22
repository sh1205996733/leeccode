package binarytree

import (
	"container/list"
)

// 二叉搜索树中第K小的元素
// https://leetcode.cn/problems/kth-smallest-element-in-a-bst/description/
func kthSmallest(root *TreeNode, k int) int {
	stack := list.New()
	node := root
	for node != nil || stack.Len() > 0 {
		if node != nil {
			stack.PushBack(node)
			node = node.Left
		} else {
			node = stack.Remove(stack.Back()).(*TreeNode)
			//
			k--
			if k == 0 {
				return node.Val
			}
			node = node.Right
		}
	}
	return 0
}
