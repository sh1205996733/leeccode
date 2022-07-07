package main

import (
	"container/list"
	"fmt"
)

func main() {
	nums := []int{-10, -3, 0, 5, 9}
	fmt.Printf("%+v", kthSmallest(nums))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func kthSmallest(root *TreeNode, k int) int {
	stack := list.New()
	node := root
	for node != nil || stack.Len() > 0 {
		if node != nil {
			stack.PushBack(node)
			node = node.Left
		} else {
			node = stack.Remove(stack.Back()).(*TreeNode)
			//
			k--
			if k == 0 {
				return node.Val
			}
			node = node.Right
		}
	}
	return 0
}
