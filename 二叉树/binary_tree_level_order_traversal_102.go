package binarytree

// 二叉树的层序遍历
// https://leetcode.cn/problems/binary-tree-level-order-traversal/description/

// 方法一：迭代
func levelOrder(root *TreeNode) (arr [][]int) {
	var ans [][]int
	var queue []*TreeNode
	queue = append(queue, root)
	level := len(queue)
	var ret []int
	for len(queue) > 0 {
		node := queue[0]
		ret = append(ret, node.Val)
		queue = queue[1:]
		if node.Left != nil {
			queue = append(queue, node.Left)
		}
		if node.Right != nil {
			queue = append(queue, node.Right)
		}
		level--
		if level == 0 {
			level = len(queue)
			ans = append(ans, ret)
			ret = nil
		}
	}
	return ans
}
