package binarytree

// 二叉树展开为链表
// https://leetcode.cn/problems/flatten-binary-tree-to-linked-list/description/
// 方法一：先序遍历
func flatten0(root *TreeNode) {
	if root == nil {
		return
	}
	node := root
	stack := []*TreeNode{node}
	tail := new(TreeNode)
	for len(stack) > 0 {
		node = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if node.Right != nil {
			stack = append(stack, node.Right)
		}
		if node.Left != nil {
			stack = append(stack, node.Left)
		}
		node.Left = nil
		tail.Right = node
		tail = node
	}
}

// 递归
func flatten(root *TreeNode) {
	dfs(root)
}

func dfs(node *TreeNode) *TreeNode {
	if node == nil || node.Left == nil && node.Right == nil {
		return node
	}
	left := dfs(node.Left)
	right := dfs(node.Right)
	if left != nil {
		left.Right = node.Right //左边连接右边
		node.Right = node.Left  //右边改成左边
		node.Left = nil
	}
	if right != nil {
		return right
	}
	return left
}
