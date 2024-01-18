package binarytree

import (
	"fmt"
	"testing"
)

func TestInorderTraversal(t *testing.T) {
	root := &TreeNode{Val: 10, Left: &TreeNode{Val: 5}, Right: &TreeNode{Val: 25}}
	fmt.Println(inorderTraversal(root))
}

func TestIsValidBST(t *testing.T) {
	root := &TreeNode{Val: 10, Left: &TreeNode{Val: 5}, Right: &TreeNode{Val: 25}}
	fmt.Println(isValidBST(root))
}

func TestRecoverTree(t *testing.T) {
	root := &TreeNode{Val: 10, Left: &TreeNode{Val: 5}, Right: &TreeNode{Val: 25}}
	recoverTree(root)
}
