package main

import "container/list"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {

}

//func preorderTraversal(root *TreeNode) []int {
//	if root == nil {
//		return []int{}
//	}
//	return  append(append([]int{root.Val}, preorderTraversal(root.Left)...), preorderTraversal(root.Right)...)
//}
func preorderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	return append(append([]int{root.Val}, preorderTraversal(root.Left)...), preorderTraversal(root.Right)...)
}

//迭代 内耗大
func preorderTraversal(root *TreeNode) (arr []int) {
	if root == nil {
		return []int{}
	}
	stack := list.New()
	stack.PushBack(root)
	for stack.Len() > 0 {
		top := stack.Remove(stack.Back()).(*TreeNode)
		arr = append(arr, top.Val)
		if top.Right != nil {
			stack.PushBack(top.Right)
		}
		if top.Left != nil {
			stack.PushBack(top.Left)
		}
	}
	return arr
}

//迭代 内耗小
func preorderTraversal(root *TreeNode) (arr []int) {
	if root == nil {
		return []int{}
	}
	var stack []*TreeNode
	stack = append(stack, root)
	for len(stack) > 0 {
		top := stack[len(stack)-1]
		arr = append(arr, top.Val)
		stack = stack[:len(stack)-1]
		if top.Right != nil {
			stack = append(stack, top.Right)
		}
		if top.Left != nil {
			stack = append(stack, top.Left)
		}
	}
	return arr
}
