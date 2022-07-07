package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//中序遍历 inorder = [9,3,15,20,7]
//后序遍历 postorder = [9,15,7,20,3]
func buildTree(inorder []int, postorder []int) *TreeNode {
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
