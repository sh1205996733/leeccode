package binarytree

// 平衡二叉树
// https://leetcode.cn/problems/balanced-binary-tree/description/
func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}
	h := height(root.Left) - height(root.Right)
	return h <= 1 && h >= -1 && isBalanced(root.Left) && isBalanced(root.Right)
}

func height(node *TreeNode) int {
	if node == nil {
		return 0
	}
	return max(height(node.Left), height(node.Right)) + 1
}
