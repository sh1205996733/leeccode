package binarytree

// 删除二叉搜索树中的节点
// https://leetcode.cn/problems/delete-node-in-a-bst/description/

func successor(root *TreeNode) int {
	root = root.Right
	for root.Left != nil {
		root = root.Left
	}
	return root.Val
}
func predecessor(root *TreeNode) int {
	root = root.Left
	for root.Right != nil {
		root = root.Right
	}
	return root.Val
}

func deleteNode0(root *TreeNode, key int) *TreeNode {
	if root == nil {
		return root
	}
	if key > root.Val {
		root.Right = deleteNode(root.Right, key)
	} else if key < root.Val {
		root.Left = deleteNode(root.Left, key)
	} else {
		if root.Left == nil && root.Right == nil {
			return nil
		} else if root.Right != nil {
			root.Val = successor(root)
			root.Right = deleteNode(root.Right, root.Val)
		} else {
			root.Val = predecessor(root)
			root.Left = deleteNode(root.Left, root.Val)
		}
	}
	return root
}

func deleteNode(root *TreeNode, key int) *TreeNode {
	var cur, curParent *TreeNode = root, nil
	for cur != nil && cur.Val != key {
		curParent = cur
		if cur.Val > key {
			cur = cur.Left
		} else {
			cur = cur.Right
		}
	}
	if cur == nil {
		return root
	}
	if cur.Left == nil && cur.Right == nil {
		cur = nil
	} else if root.Right != nil {
		cur = cur.Left
	} else if root.Left != nil {
		cur = cur.Right
	} else {
		successor, successorParent := cur.Right, cur
		for successor.Left != nil {
			successorParent = successor
			successor = successor.Left
		}
		if successorParent.Val == cur.Val {
			successorParent.Right = successor.Right
		} else {
			successorParent.Left = successor.Right
		}
		successor.Right = cur.Right
		successor.Left = cur.Left
		cur = successor
	}
	if curParent == nil {
		return cur
	}
	if curParent.Left != nil && curParent.Left.Val == key {
		curParent.Left = cur
	} else {
		curParent.Right = cur
	}
	return root
}
