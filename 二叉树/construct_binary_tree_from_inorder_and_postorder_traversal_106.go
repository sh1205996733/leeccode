package binarytree

// 从中序与后序遍历序列构造二叉树
// https://leetcode.cn/problems/construct-binary-tree-from-inorder-and-postorder-traversal/description/

// 方法一:递归
func buildTree2(inorder []int, postorder []int) *TreeNode {
	if len(postorder) == 0 || len(inorder) == 0 {
		return nil
	}
	lastIndex := len(postorder) - 1
	root := &TreeNode{Val: postorder[lastIndex]}
	for key, value := range inorder {
		if value == postorder[lastIndex] {
			root.Left = buildTree(inorder[:key], postorder[:key])
			root.Right = buildTree(inorder[key+1:], postorder[key:len(postorder)-1])
		}
	}
	return root
}
