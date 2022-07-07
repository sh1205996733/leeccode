package main

import (
	"container/list"
	"fmt"
)

func main() {
	nums := []int{-10, -3, 0, 5, 9}
	fmt.Printf("%+v", rangeSumBST(nums))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//func rangeSumBST(root *TreeNode, low int, high int) int {
//	stack := list.New()
//	node := root
//	ret := 0
//	for node != nil || stack.Len() > 0 {
//		if node != nil {
//			stack.PushBack(node)
//			node = node.Left
//		}else {
//			node = stack.Remove(stack.Back()).(*TreeNode)
//			//
//			if node.Val>=low && node.Val <= high {
//				ret += node.Val
//			}
//			node = node.Right
//		}
//	}
//	return ret
//}
func rangeSumBST(root *TreeNode, low int, high int) int {
	ret := sumBST(root, low, high)
	return ret
}
func sumBST(root *TreeNode, low int, high int) int {
	if root == nil {
		return 0
	}
	ret := sumBST(root.Left, low, high)
	if root.Val >= low && root.Val <= high {
		ret += root.Val
	}
	return ret + sumBST(root.Right, low, high)
}
