package linkedlist

import binarytree "leetcode/二叉树"

// https://leetcode.cn/problems/populating-next-right-pointers-in-each-node
// 填充每个节点的下一个右侧节点指针

// 方法一 层序遍历
func connect0(root *binarytree.TreeNode) *binarytree.TreeNode {
	if root == nil {
		return root
	}
	queue := make([]*binarytree.TreeNode, 0)
	queue = append(queue, root)
	levelCount := 1
	for len(queue) > 0 {
		n := queue[0]
		queue = queue[1:]
		levelCount--
		if n.Left != nil {
			queue = append(queue, n.Left)
		}
		if n.Right != nil {
			queue = append(queue, n.Right)
		}
		if levelCount == 0 {
			levelCount = len(queue)
			n.Next = nil
		} else {
			n.Next = queue[0]
		}
	}
	return root
}

// 只使用常量级额外空间
func connect(root *binarytree.TreeNode) *binarytree.TreeNode {
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
