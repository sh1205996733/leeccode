package binarytree

import (
	"fmt"
	"github.com/bmizerany/assert"
	"leetcode/二叉树/printer"
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

func TestIsSymmetric(t *testing.T) {
	nums := []any{1, 2, 2, 3, 4, 4, 3}
	//nums := []any{1, 2, 2, nil, 3, nil, 3}
	root := MakePerfectBinaryTree(nums)
	fmt.Println(isSymmetric(root))
}

func TestLevelOrder(t *testing.T) {
	nums := []any{1, 2, 2, 3, 4, 4, 3}
	root := MakePerfectBinaryTree(nums)
	fmt.Println(levelOrder(root))
}

func TestLevelOrderBottom(t *testing.T) {
	nums := []any{1, 2, 2, 3, 4, 4, 3}
	root := MakePerfectBinaryTree(nums)
	fmt.Println(levelOrderBottom(root))
}

func TestMaxDepth(t *testing.T) {
	nums := []any{1, 2, 2, nil, 3, nil, 3, nil, nil, 1}
	root := MakePerfectBinaryTree(nums)
	fmt.Println(maxDepth(root))
}

func TestBuildTree(t *testing.T) {
	preorder := []int{3, 9, 1, 10, 20, 15, 7}
	inorder := []int{1, 9, 10, 3, 15, 20, 7}
	MakePerfectBinaryTree([]any{3, 9, 20, 1, 10, 15, 7})
	printer.Println(buildTree(preorder, inorder))
}

func TestSortedArrayToBST(t *testing.T) {
	nums := []int{-10}
	printer.Println(sortedArrayToBST(nums))
}

func TestIsBalanced(t *testing.T) {
	nums := []any{3, 9, 20, nil, nil, 15, 7}
	root := MakePerfectBinaryTree(nums)
	fmt.Println(isBalanced(root))
}

func TestFlatten(t *testing.T) {
	nums := []any{1, 2, 5, 3, 4, nil, 6}
	root := MakePerfectBinaryTree(nums)
	flatten(root)
	printer.Println(root)
}

func TestWidthOfBinaryTree(t *testing.T) {
	//nums := []any{1, 3, 2, 5, 3, nil, 9} //4
	nums := []any{1, 3, 2, 5, nil, nil, 9, 6, nil, nil, nil, nil, nil, 7, nil} // 7
	//nums := []any{1, 3, 2, 5} // 2
	root := MakePerfectBinaryTree(nums)
	assert.Equal(t, widthOfBinaryTree(root), 2)
}

func TestConstructFromPrePost(t *testing.T) {
	preorder := []int{1, 2, 4, 5, 3, 6, 7}
	postorder := []int{4, 5, 2, 6, 7, 3, 1}
	printer.Println(constructFromPrePost(preorder, postorder))
}
