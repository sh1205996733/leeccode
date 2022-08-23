package linkedlist

// https://leetcode.cn/problems/flatten-a-multilevel-doubly-linked-list/
// https://leetcode.cn/problems/Qv1Da2/
// 扁平化多级双向链表

// 方法一：递归 对child节点进行flatten操作，再遍历child节点找到尾节点
// 时间复杂度：O(N^2)
// 空间复杂度：O(n)
func flatten0(root *Node) *Node {
	for node := root; node != nil; node = node.Next {
		if node.Child != nil {
			next := node.Next
			child := flatten(node.Child)
			node.Next = child
			child.Prev = node
			node.Child = nil
			for ; node.Next != nil; node = node.Next { //找尾节点
			}
			if next != nil {
				node.Next = next
				next.Prev = node
			}
		}
	}
	return root
}

// 方法二：递归优化 对child节点进行flatten操作，直接返回尾节点
// 时间复杂度：O(N)
// 空间复杂度：O(n)
func flatten(root *Node) *Node {
	dfs(root)
	return root
}

// 返回尾节点
func dfs(root *Node) *Node {
	last := root
	for node := root; node != nil; node = node.Next {
		if node.Child != nil {
			next := node.Next
			childLast := dfs(node.Child)
			node.Next = node.Child
			node.Child.Prev = node
			node.Child = nil
			if next != nil {
				childLast.Next = next
				next.Prev = childLast
			}
			node = childLast
		}
		last = node
	}
	return last
}

// 方法三：迭代一段一段拼接
// 时间复杂度：O(N)
// 空间复杂度：O(1)
func flatten1(root *Node) *Node {
	for node := root; node != nil; node = node.Next {
		if node.Child != nil {
			next := node.Next
			node.Next = node.Child
			node.Child.Prev = node
			node.Child = nil
			last := node
			for ; last.Next != nil; last = last.Next {
			}
			if next != nil {
				last.Next = next
				next.Prev = last
			}
		}
	}
	return root
}
