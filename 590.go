package main

import "fmt"

type Node struct {
	Val      int
	Children []*Node
}

var a []int

func postorder1(root *Node) []int {
	a = []int{}
	dfs(root)
	return a
}
func dfs(root *Node) {
	if root != nil {
		for _, v := range root.Children {
			dfs(v)
		}
		a = append(a, root.Val)
	}
}

func postorder(root *Node) []int {
	if root == nil {
		return []int{}
	}
	var ret []int
	var stack []*Node
	var pres []*Node
	node := root
	stack = append(stack, node)
	for len(stack) > 0 {
		node := stack[len(stack)-1]
		if len(node.Children) == 0 || (len(pres) > 0 && node == pres[len(pres)-1]) {
			if len(pres) > 0 && node == pres[len(pres)-1] {
				pres = pres[:len(pres)-1]
			}
			ret = append(ret, node.Val)
			stack = stack[:len(stack)-1]
		} else {
			for i := len(node.Children) - 1; i >= 0; i-- {
				stack = append(stack, node.Children[i])
			}
			pres = append(pres, node)
		}
	}
	return ret
}

func main() {
	postorder(nil)
	fmt.Println("---------")
}
