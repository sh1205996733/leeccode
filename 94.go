package main

import "container/list"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {

}
func inorderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	return append(append(inorderTraversal(root.Left), root.Val), inorderTraversal(root.Right)...)
}

//迭代 内耗大
func inorderTraversal(root *TreeNode) (arr []int) {
	if root == nil {
		return []int{}
	}
	stack := list.New()
	node := root
	for node != nil || stack.Len() > 0 {
		if node != nil {
			stack.PushBack(node)
			node = node.Left
		} else {
			node = stack.Remove(stack.Back()).(*TreeNode)
			arr = append(arr, node.Val)
			node = node.Right
		}
	}
	return arr
}

//迭代 内耗小
func inorderTraversal(root *TreeNode) (arr []int) {
	if root == nil {
		return []int{}
	}
	var stack []*TreeNode
	node := root
	for node != nil || len(stack) > 0 {
		if node != nil {
			stack = append(stack, node)
			node = node.Left
		} else {
			node = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			arr = append(arr, node.Val)
			node = node.Right
		}
	}
	return arr
}
