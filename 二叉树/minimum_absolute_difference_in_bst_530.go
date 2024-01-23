package binarytree

// 二叉搜索树的最小绝对差
// https://leetcode.cn/problems/minimum-absolute-difference-in-bst/description/
// 二叉搜索树节点最小距离
//  https://leetcode-cn.com/problems/minimum-distance-between-bst-nodes/
import (
	"container/list"
	"math"
)

// 非递归-栈实现
func getMinimumDifference0(root *TreeNode) int {
	min := math.MaxInt64
	var minVal = -1
	stack := list.New()
	node := root
	for node != nil || stack.Len() > 0 {
		if node != nil {
			stack.PushBack(node)
			node = node.Left
		} else {
			node = stack.Remove(stack.Back()).(*TreeNode)
			abs := node.Val - minVal
			if minVal != -1 {
				if abs < min {
					min = abs
				}
			}
			minVal = node.Val
			node = node.Right
		}
	}
	return min
}

var min = math.MaxInt64
var minVal = -1

func getMinimumDifference(root *TreeNode) int {
	minimumDifference(root)
	return min
}

// 递归实现
func minimumDifference(root *TreeNode) {
	if root == nil {
		return
	}
	minimumDifference(root.Left)
	abs := root.Val - minVal
	if minVal != -1 {
		if abs < min {
			min = abs
		}
	}
	minVal = root.Val
	minimumDifference(root.Right)
}
