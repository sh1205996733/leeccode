package binarytree

// 二叉树的最大深度
// https://leetcode.cn/problems/maximum-depth-of-binary-tree/description/
// 方法一:递归
func maxDepth0(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return max(maxDepth(root.Left), maxDepth(root.Right)) + 1
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

// 方法一:层序遍历
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	var queue []*TreeNode
	height := 0
	levelCount := 1
	queue = append(queue, root)
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		levelCount--
		if node.Left != nil {
			queue = append(queue, node.Left)
		}
		if node.Right != nil {
			queue = append(queue, node.Right)
		}
		if levelCount == 0 {
			height++
			levelCount = len(queue)
		}
	}
	return height
}
