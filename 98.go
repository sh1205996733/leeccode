package main

import (
	"container/list"
	"math"
)

func main() {
	isValidBST(nil)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//var stack = list.New()
//func isValidBST(root *TreeNode) bool {
//	node := root
//	min := math.MinInt64
//	for stack.Len() > 0 {
//		if node != nil {
//			stack.PushBack(node)
//			node = node.Left
//		} else {//左边为空
//			node = stack.Remove(stack.Back()).(*TreeNode)
//			// 访问节点
//			if node.Val <= min {
//				return false
//			}
//			min = node.Val
//			node = node.Right
//		}
//	}
//	return true
//}

// 递归解法
//func isValidBST(root *TreeNode) bool {
//	// 初始上下界为max和min
//	return isBST(root, math.MinInt64, math.MaxInt64)
//}
//
//// root：根
//// lower：下界
//// upper：上界
//func isBST(root *TreeNode, lower, upper int) bool {
//	// 空树是二叉树
//	if root == nil {
//		return true
//	}
//	// 二叉树定义：根节点的值要大于左节点且小于右节点
//	// 不满足定义就不是二叉树
//	if root.Val <= lower || root.Val >= upper {
//		return false
//	}
//
//	// 判断左子树是不是二叉树
//	// 因为左子树里所有节点的值都要小于根节点
//	// 所以上界 upper 改为 root.Val
//	isLeft := isBST(root.Left, lower, root.Val)
//	// 判断右子树是不是二叉树
//	// 因为右子树里所有节点的值要大于根节点
//	// 所以下界 lower 改为 root.Val
//	isRight := isBST(root.Right, root.Val, upper)
//
//	// 左右子树都是二叉树，才是二叉树
//	return isLeft && isRight
//}
