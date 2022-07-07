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
func postorderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	return append(append(postorderTraversal(root.Left), postorderTraversal(root.Right)...), root.Val)
}

//迭代 内耗大
func postorderTraversal(root *TreeNode) (arr []int) {
	if root == nil {
		return []int{}
	}
	stack := list.New()
	node := root
	stack.PushBack(node)
	for node != nil || stack.Len() > 0 {

	}
	return arr
}

//迭代 内耗小
func postorderTraversal(root *TreeNode) (arr []int) {
	if root == nil {
		return []int{}
	}
	var pre *TreeNode
	var stack []*TreeNode
	stack = append(stack, root)
	for len(stack) > 0 {
		top := stack[len(stack)-1]
		if (top.Left == nil && top.Right == nil) || (pre != nil && (pre == top.Left || pre == top.Right)) {
			arr = append(arr, top.Val)
			pre = top
			stack = stack[:len(stack)-1]
		} else {
			if top.Right != nil {
				stack = append(stack, top.Right)
			}
			if top.Left != nil {
				stack = append(stack, top.Left)
			}
		}
	}
	return arr
}
