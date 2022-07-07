package main

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return math.Abs(height(root.Left)-height(root.Right)) <= 1 && isBalanced(root.Left) && isBalanced(root.Right)
}
func height(root *TreeNode) float64 {
	if root == nil {
		return 0
	}
	return math.Max(height(root.Left), height(root.Right)) + 1
}
