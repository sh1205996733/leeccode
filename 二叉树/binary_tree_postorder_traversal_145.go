package binarytree

// 二叉树的后序遍历
// https://leetcode.cn/problems/binary-tree-postorder-traversal/description/

func postorderTraversal0(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	return append(append(postorderTraversal(root.Left), postorderTraversal(root.Right)...), root.Val)
}

// 迭代 内耗小
func postorderTraversal(root *TreeNode) (arr []int) {
	if root == nil {
		return []int{}
	}
	var pre *TreeNode
	var stack []*TreeNode
	stack = append(stack, root)
	for len(stack) > 0 {
		top := stack[len(stack)-1]
		if (top.Left == nil && top.Right == nil) || (pre != nil && (pre == top.Left || pre == top.Right)) {
			arr = append(arr, top.Val)
			pre = top
			stack = stack[:len(stack)-1]
		} else {
			if top.Right != nil {
				stack = append(stack, top.Right)
			}
			if top.Left != nil {
				stack = append(stack, top.Left)
			}
		}
	}
	return arr
}
