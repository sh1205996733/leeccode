package binarytree

// N 叉树的前序遍历
// https://leetcode.cn/problems/n-ary-tree-preorder-traversal/description/

func preorderN0(root *Node) []int {
	if root == nil {
		return nil
	}
	ans := []int{root.Val}
	for _, v := range root.Children {
		ret := preorderN0(v)
		if len(ret) > 0 {
			ans = append(ans, ret...)
		}
	}
	return ans
}

func preorderN(root *Node) []int {
	if root == nil {
		return []int{}
	}
	var ret []int
	var stack []*Node
	node := root
	stack = append(stack, node)
	for len(stack) > 0 {
		node := stack[len(stack)-1]
		ret = append(ret, node.Val)
		stack = stack[:len(stack)-1]
		if node != nil {
			for i := len(node.Children) - 1; i >= 0; i-- {
				stack = append(stack, node.Children[i])
			}
		}
	}
	return ret
}
