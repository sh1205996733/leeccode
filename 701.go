package main

func main() {
	insertIntoBST(nil, 3)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//非递归
func insertIntoBST(root *TreeNode, val int) *TreeNode {
	newNode := &TreeNode{Val: val}
	if root == nil {
		return newNode
	}

	node := root
	for node != nil {
		if val > node.Val {
			if node.Right == nil {
				node.Right = newNode
				break
			}
			node = node.Right
		} else {
			if node.Left == nil {
				node.Left = newNode
				break
			}
			node = node.Left
		}
	}
	return root
}

//递归
//func insertIntoBST(root *TreeNode, val int) *TreeNode {
//	for root != nil {
//		if val > root.Val {
//			root.Right = insertIntoBST(root.Right,val)
//		} else if val < root.Val {
//			root.Left = insertIntoBST(root.Left,val)
//		} else {
//			root.Val = val
//		}
//		return root
//	}
//	return &TreeNode{Val:val}
//}
