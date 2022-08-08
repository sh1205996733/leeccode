package linkedlist

import "leetcode/二叉树"

// https://leetcode.cn/problems/binode-lcci/
// 二叉搜索树转换为单向链表

// 方法一 迭代 中序遍历，把每个节点穿起来
func convertBiNode(root *binarytree.TreeNode) *binarytree.TreeNode {
	var stack []*binarytree.TreeNode
	newHead := &binarytree.TreeNode{0, nil, nil}
	node, prev := root, newHead
	for node != nil || len(stack) > 0 {
		if node != nil {
			stack = append(stack, node)
			node = node.Left
		} else {
			node = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			node.Left = nil
			prev.Right = node
			prev = node
			node = node.Right
		}
	}
	return newHead.Right
}

// 方法二 递归 中序遍历，把每个节点穿起来
func convertBiNode1(root *binarytree.TreeNode) *binarytree.TreeNode {
	newHead := &binarytree.TreeNode{0, nil, nil}
	convert(root, newHead)
	return newHead.Right
}

func convert(root, prev *binarytree.TreeNode) *binarytree.TreeNode {
	if root == nil {
		return prev
	}
	prev = convert(root.Left, prev)
	root.Left = nil
	prev.Right = root
	prev = root
	prev = convert(root.Right, prev)
	return prev
}
