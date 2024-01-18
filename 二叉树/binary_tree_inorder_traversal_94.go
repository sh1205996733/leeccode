package binarytree

// 二叉树的中序遍历
// https://leetcode.cn/problems/binary-tree-inorder-traversal/description/

func inorderTraversal0(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	return append(append(inorderTraversal(root.Left), root.Val), inorderTraversal(root.Right)...)
}

// 迭代
func inorderTraversal(root *TreeNode) (arr []int) {
	if root == nil {
		return []int{}
	}
	var stack []*TreeNode
	node := root
	for node != nil || len(stack) > 0 {
		if node != nil {
			stack = append(stack, node)
			node = node.Left
		} else {
			node = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			arr = append(arr, node.Val)
			node = node.Right
		}
	}
	return arr
}
