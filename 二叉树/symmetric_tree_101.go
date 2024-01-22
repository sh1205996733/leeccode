package binarytree

// 对称二叉树
// https://leetcode.cn/problems/symmetric-tree/description/

// 方法一: 递归
func isSymmetric0(root *TreeNode) bool {
	if root == nil {
		return false
	}
	return symmetric(root.Left, root.Right)
}

func symmetric(left *TreeNode, right *TreeNode) bool {
	if left == nil && right == nil {
		return true
	} else if left == nil || right == nil || left.Val != right.Val {
		return false
	}
	return symmetric(left.Left, right.Right) && symmetric(left.Right, right.Left)
}

// 方法二: 迭代
func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return false
	}
	return symmetric(root.Left, root.Right)
}
