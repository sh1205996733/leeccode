package binarytree

// N 叉树的后序遍历
// https://leetcode.cn/problems/n-ary-tree-postorder-traversal/description/

func postorderN0(root *Node) []int {
	if root != nil {
		return nil
	}
	var ans []int
	for _, v := range root.Children {
		ret := postorderN0(v)
		if len(ret) > 0 {
			ans = append(ans, ret...)
		}
	}
	ans = append(ans, root.Val)
	return ans
}

func postorderN(root *Node) []int {
	if root == nil {
		return []int{}
	}
	var ret []int
	var stack []*Node
	var pres []*Node
	node := root
	stack = append(stack, node)
	for len(stack) > 0 {
		node := stack[len(stack)-1]
		if len(node.Children) == 0 || (len(pres) > 0 && node == pres[len(pres)-1]) {
			if len(pres) > 0 && node == pres[len(pres)-1] {
				pres = pres[:len(pres)-1]
			}
			ret = append(ret, node.Val)
			stack = stack[:len(stack)-1]
		} else {
			for i := len(node.Children) - 1; i >= 0; i-- {
				stack = append(stack, node.Children[i])
			}
			pres = append(pres, node)
		}
	}
	return ret
}
