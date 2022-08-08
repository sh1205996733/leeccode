package linkedlist

import binarytree "leetcode/二叉树"

// https://leetcode.cn/problems/populating-next-right-pointers-in-each-node-ii
// 填充每个节点的下一个右侧节点指针

// 只使用常量级额外空间
func connectII(root *binarytree.TreeNode) *binarytree.TreeNode {
	if root == nil {
		return root
	}
	prev := root
	for head := root.Left; head != nil; head = head.Left { //	每一层最左边第一个数,nil就结束
		for ; prev != nil; prev = prev.Next {

			prev.Left.Next = prev.Right
			if prev.Next != nil {
				prev.Right.Next = prev.Next.Left
			}
		}
		prev = head
	}
	return root
}
