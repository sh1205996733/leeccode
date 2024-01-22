package binarytree

// 二叉树的前序遍历
// https://leetcode.cn/problems/binary-tree-preorder-traversal/description/

func preorderTraversal0(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	return append(append([]int{root.Val}, preorderTraversal(root.Left)...), preorderTraversal(root.Right)...)
}

// 迭代 内耗小
func preorderTraversal(root *TreeNode) (arr []int) {
	if root == nil {
		return []int{}
	}
	var stack []*TreeNode
	stack = append(stack, root)
	for len(stack) > 0 {
		top := stack[len(stack)-1]
		arr = append(arr, top.Val)
		stack = stack[:len(stack)-1]
		if top.Right != nil {
			stack = append(stack, top.Right)
		}
		if top.Left != nil {
			stack = append(stack, top.Left)
		}
	}
	return arr
}
