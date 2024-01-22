package binarytree

// 二叉树的层序遍历 II
// https://leetcode.cn/problems/binary-tree-level-order-traversal-ii/description/

// 方法一：迭代
func levelOrderBottom(root *TreeNode) (ret [][]int) {
	if root == nil {
		return [][]int{}
	}
	var queue []*TreeNode //模拟队列
	node := root
	queue = append(queue, node)
	levelCount := 1
	var a []int
	for len(queue) > 0 {
		node = queue[0]
		queue = queue[1:]
		a = append(a, node.Val)
		levelCount--
		if node.Left != nil {
			queue = append(queue, node.Left)
		}
		if node.Right != nil {
			queue = append(queue, node.Right)
		}
		if levelCount == 0 {
			levelCount = len(queue)
			ret = append(ret, a)
			a = make([]int, 0)
		}
	}
	start, end := 0, len(ret)-1
	for start < end {
		ret[start], ret[end] = ret[end], ret[start]
		start++
		end--
	}
	return ret
}
