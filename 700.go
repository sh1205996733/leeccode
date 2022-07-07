package main

func main() {
	deleteNode(nil, 3)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func (p *TreeNode) Root() interface{} {
	return p
}
func (p *TreeNode) LNode() interface{} {
	if p.Left == nil {
		return nil
	}
	return p.Left
}
func (p *TreeNode) RNode() interface{} {
	if p.Right == nil {
		return nil
	}
	return p.Right
}
func (p *TreeNode) ToString() interface{} {
	return p.Val
}
func (p *TreeNode) ColorOf() bool {
	return false
}
func searchBST1(root *TreeNode, val int) *TreeNode {
	if root == nil || root.Val == val {
		return root
	}
	if root.Val < val {
		return searchBST(root.Right, val)
	} else {
		return searchBST(root.Left, val)
	}
}

//非递归
func searchBST(root *TreeNode, val int) *TreeNode {
	for root != nil && root.Val != val {
		if root.Val < val {
			root = root.Right
		} else {
			root = root.Left
		}
	}
	return root
}
