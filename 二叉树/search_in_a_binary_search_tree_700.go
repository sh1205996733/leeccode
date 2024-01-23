package binarytree

// 二叉搜索树中的搜索
//https://leetcode.cn/problems/search-in-a-binary-search-tree/description/

func searchBST1(root *TreeNode, val int) *TreeNode {
	if root == nil || root.Val == val {
		return root
	}
	if root.Val < val {
		return searchBST(root.Right, val)
	} else {
		return searchBST(root.Left, val)
	}
}

// 非递归
func searchBST(root *TreeNode, val int) *TreeNode {
	node := root
	for node != nil && node.Val != val {
		if node.Val < val {
			node = node.Right
		} else {
			node = node.Left
		}
	}
	return node
}
