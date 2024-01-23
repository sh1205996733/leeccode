package binarytree

// 根据前序和后序遍历构造二叉树
// https://leetcode.cn/problems/construct-binary-tree-from-preorder-and-postorder-traversal/description/

func constructFromPrePost(preorder []int, postorder []int) *TreeNode {
	if len(postorder) == 0 {
		return nil
	}
	if len(postorder) == 1 {
		return &TreeNode{Val: preorder[0]}
	}
	root := &TreeNode{Val: preorder[0]}
	var idx int
	for i := range postorder { //找左节点位置
		if postorder[i] == preorder[1] {
			idx = i
			break
		}
	}
	root.Left = constructFromPrePost(preorder[1:idx+2], postorder[0:idx+1])
	root.Right = constructFromPrePost(preorder[idx+2:], postorder[idx+1:len(postorder)-1])
	return root
}
