package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//前序遍历 preorder = [3,9,20,15,7]
//中序遍历 inorder = [9,3,15,20,7]
func buildTree1(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}
	root := &TreeNode{Val: preorder[0]}
	for key, value := range inorder {
		if value == preorder[0] {
			root.Left = buildTree(preorder[1:len(inorder[:key])+1], inorder[:key])
			root.Right = buildTree(preorder[len(inorder[:key])+1:], inorder[key+1:])
		}
	}
	return root
}

//迭代
/**
v 是u的左儿子。这是因为在遍历到u之后，下一个遍历的节点就是u的左儿子，即v；
u 没有左儿子，并且v是u的某个祖先节点（或者u本身）的右儿子。如果u没有左儿子，那么下一个遍历的节点就是u的右儿子。如果u没有右儿子，
我们就会向上回溯，直到遇到第一个有右儿子（且u不在它的右儿子的子树中）的节点 ua​，那么v就是ua的右儿子。
*/
func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}
	root := &TreeNode{Val: preorder[0]}
	stack := []*TreeNode{root}
	var inorderIndex int
	for i := 1; i < len(inorder); i++ {
		preorderVal := preorder[i]
		node := stack[len(stack)-1]
		if node.Val != inorder[inorderIndex] {
			node.Left = &TreeNode{preorderVal, nil, nil}
			stack = append(stack, node.Left)
		} else {
			for len(stack) != 0 && stack[len(stack)-1].Val == inorder[inorderIndex] {
				node = stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				inorderIndex++
			}
			node.Right = &TreeNode{preorderVal, nil, nil}
			stack = append(stack, node.Right)
		}
	}
	return root
}
