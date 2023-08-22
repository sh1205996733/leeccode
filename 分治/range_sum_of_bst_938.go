package divide_conquer

import binarytree "leetcode/二叉树"

// https://leetcode.cn/problems/range-sum-of-bst/

func rangeSumBST(root *binarytree.TreeNode, low int, high int) int {
	ret := sumBST(root, low, high)
	return ret
}
func sumBST(root *binarytree.TreeNode, low int, high int) int {
	if root == nil {
		return 0
	}
	ret := sumBST(root.Left, low, high)
	if root.Val >= low && root.Val <= high {
		ret += root.Val
	}
	return ret + sumBST(root.Right, low, high)
}

// 迭代
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
