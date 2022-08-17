package linkedlist

import binarytree "leetcode/二叉树"

// https://leetcode.cn/problems/convert-sorted-list-to-binary-search-tree/
// 有序链表转换二叉搜索树

// 方法一 分治 想象成一条绳，提起中点作为根节点，分出左右两部分，再提起各自的中点作为根节点……分治下去，这根绳就成了BST的模样
// 时间复杂度：O(nlogn)
// 空间复杂度：O(logn)
func sortedListToBST0(head *ListNode) *binarytree.TreeNode {
	if head == nil {
		return nil
	}
	mid, fast := head, head
	var prev *ListNode
	for fast != nil && fast.Next != nil {
		prev = mid
		mid = mid.Next
		fast = fast.Next.Next
	}
	root := &binarytree.TreeNode{Val: mid.Val}
	if prev != nil { //mid不再重复计算
		prev.Next = nil
		root.Left = sortedListToBST(head)
	}
	root.Right = sortedListToBST(mid.Next)
	return root
}

// 方法二 中序遍历策略带来的优化
// 时间复杂度：O(n)
// 空间复杂度：O(logn)
var h *ListNode

func sortedListToBST(head *ListNode) *binarytree.TreeNode {
	if head == nil {
		return nil
	}
	length := 0
	h = head
	for head != nil {
		length++
		head = head.Next
	}
	return buildBST(0, length-1)
}
func buildBST(start, end int) *binarytree.TreeNode {
	if start > end {
		return nil
	}
	mid := (start + end) >> 1
	left := buildBST(start, mid-1)
	root := &binarytree.TreeNode{Val: h.Val}
	h = h.Next
	root.Left = left
	root.Right = buildBST(mid+1, end)
	return root
}
