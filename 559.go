package main

import "fmt"

type Node struct {
	Val      int
	Children []*Node
}

func maxDepth(root *Node) int {
	if root == nil {
		return 0
	}
	var stack []*Node
	levelCount := 1
	height := 0
	node := root
	stack = append(stack, node)
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

//var height
//func maxDepth(root *Node) int {
//	if root == nil {
//		return 0
//	}
//	height = 1
//	dfs(root)
//	return height
//}
//func dfs(root *Node) {
//	if root != nil && len(root.Children) > 0{
//		height++
//		for _,v := range root.Children {
//			dfs(v)
//		}
//	}
//}

func main() {
	postorder(nil)
	fmt.Println("---------")
}
