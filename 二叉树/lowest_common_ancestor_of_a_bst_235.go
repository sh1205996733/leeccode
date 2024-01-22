package binarytree

// https://leetcode.cn/problems/lowest-common-ancestor-of-a-binary-search-tree/description/
// 二叉搜索树的最近公共祖先

func lowestCommonAncestor0(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	//p、q不是root
	if p.Val < root.Val && q.Val < root.Val { //p、q都在左边
		return lowestCommonAncestor(root.Left, p, q)
	} else if p.Val > root.Val && q.Val > root.Val { //p、q都在右边
		return lowestCommonAncestor(root.Right, p, q)
	} else {
		return root
	}
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	var ancestor = root
	for {
		if p.Val < ancestor.Val && q.Val < ancestor.Val {
			ancestor = ancestor.Left
		} else if p.Val > ancestor.Val && q.Val > ancestor.Val {
			ancestor = ancestor.Right
		} else {
			return ancestor
		}
	}
	return ancestor
}
