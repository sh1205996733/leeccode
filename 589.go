package main

type Node struct {
	Val      int
	Children []*Node
}

var a []int

func preorder(root *Node) []int {
	a = []int{}
	dfs(root)
	return a
}
func dfs(root *Node) {
	if root != nil {
		a = append(a, root.Val)
		for _, v := range root.Children {
			dfs(v)
		}
	}
}

func preorder(root *Node) []int {
	if root == nil {
		return []int{}
	}
	var ret []int
	var stack []*Node
	node := root
	stack = append(stack, node)
	for len(stack) > 0 {
		node := stack[len(stack)-1]
		ret = append(ret, node.Val)
		stack = stack[:len(stack)-1]
		if node != nil {
			for i := len(node.Children) - 1; i >= 0; i-- {
				stack = append(stack, node.Children[i])
			}
		}
	}
	return ret
}
