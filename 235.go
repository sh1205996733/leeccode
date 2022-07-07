package main

import "fmt"

func main() {
	nums := []int{-10, -3, 0, 5, 9}
	fmt.Printf("%+v", rangeSumBST(nums))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
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
