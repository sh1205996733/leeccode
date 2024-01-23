package binarytree

// N 叉树的最大深度
// https://leetcode.cn/problems/maximum-depth-of-n-ary-tree/description/

type Node struct {
	Val      int
	Children []*Node
}

// 层序遍历
func maxDepth1(root *Node) int {
	if root == nil {
		return 0
	}
	var stack []*Node
	levelCount := 1
	height := 0
	stack = append(stack, root)
	for len(stack) > 0 {
		node := stack[0]
		stack = stack[1:]
		levelCount--
		for i := len(node.Children) - 1; i >= 0; i-- {
			stack = append(stack, node.Children[i])
		}
		if levelCount == 0 {
			levelCount = len(stack)
			height++
		}
	}
	return height
}

func maxDepth2(root *Node) int {
	if root == nil {
		return 0
	}
	height := 1
	dfsNode(root, &height)
	return height
}
func dfsNode(root *Node, height *int) {
	if root != nil && len(root.Children) > 0 {
		*height++
		for _, v := range root.Children {
			dfsNode(v, height)
		}
	}
}
