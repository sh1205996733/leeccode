package binarytree

// N 叉树的层序遍历
// https://leetcode.cn/problems/n-ary-tree-level-order-traversal/description/

func levelOrderN(root *Node) (ans [][]int) {
	if root == nil {
		return
	}
	q := []*Node{root}
	for q != nil {
		level := []int{}
		tmp := q
		q = nil
		for _, node := range tmp {
			level = append(level, node.Val)
			q = append(q, node.Children...)
		}
		ans = append(ans, level)
	}
	return
}
