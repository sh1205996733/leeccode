package recursion

import "math"

// 124. 二叉树中的最大路径和
// https://leetcode.cn/problems/binary-tree-maximum-path-sum/
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 方法一: 递归
// 时间复杂度：O(n)
// 空间复杂度：O(n)
var maxSum = math.MinInt32

func maxPathSum(root *TreeNode) int {
	dfs(root)
	return maxSum
}
func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func dfs(root *TreeNode) int {
	if root == nil {
		return 0
	}
	left := dfs(root.Left)   //递归左子树最大路径和
	right := dfs(root.Right) //递归右子树最大路径和
	maxSum = Max(maxSum, left+root.Val+right)
	pathSum := root.Val + Max(Max(0, left), right)
	if pathSum < 0 {
		return 0
	}
	return pathSum
}
