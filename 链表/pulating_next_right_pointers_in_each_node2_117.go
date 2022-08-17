package linkedlist

import binarytree "leetcode/二叉树"

// https://leetcode.cn/problems/populating-next-right-pointers-in-each-node-ii
// 填充每个节点的下一个右侧节点指针

// 只使用常量级额外空间
func connectII(root *binarytree.TreeNode) *binarytree.TreeNode {
	parent := root
	var left *binarytree.TreeNode
	for parent != nil { //	每一层最左边第一个数,nil就结束
		for parent != nil {
			left = getLeft(parent)
			if left == nil {
				parent = parent.Next
			} else {
				break
			}
		}
		for ; parent != nil; parent = parent.Next {
			if parent.Left != nil {
				if parent.Right != nil {
					parent.Left.Next = parent.Right
				} else if parent.Next != nil {
					parent.Left.Next = getLeft(parent.Next)
				}
			}
			if parent.Right != nil && parent.Next != nil {
				parent.Right.Next = getLeft(parent.Next)
			}
		}
		parent = left
	}
	return root
}

func getLeft(root *binarytree.TreeNode) *binarytree.TreeNode {
	node := root.Left
	if node == nil {
		node = root.Right
	}
	if node == nil && root.Next != nil {
		return getLeft(root.Next)
	}
	return node
}
