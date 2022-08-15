package linkedlist

//https://leetcode.cn/problems/next-greater-node-in-linked-list/
// 链表中的下一个更大节点

// 方法一：暴力法，1.找第一个比他最的值，2.然后开始从头开始遍历 直到遇到最大值，再重复1,2直到结束
// 时间复杂度：O(N^2)
// 空间复杂度：O(1)
func nextLargerNodes0(head *ListNode) []int {
	var ret []int
	for node := head; node != nil; node = node.Next {
		max := node.Val
		for cur := node.Next; cur != nil; cur = cur.Next {
			if cur.Val > max {
				max = cur.Val
				break
			}
		}
		if max == node.Val {
			ret = append(ret, 0)
		} else {
			ret = append(ret, max)
		}
	}
	return ret
}

// 方法一：栈，先降序入栈，直到node比栈顶大，然后出栈
// 时间复杂度：O(N^2)
// 空间复杂度：O(1)
func nextLargerNodes(head *ListNode) []int {
	var ret []int
	var stack []int
	cur := 0
	for node := head; node != nil; node = node.Next {
		for len(stack) != 0 && stack[len(stack)-1] < node.Val {
			stack = stack[0 : len(stack)-1]
			for ret[cur] != 0 {
				cur--
			}
			ret[cur] = node.Val
		}
		stack = append(stack, node.Val)
		ret = append(ret, 0)
		cur = len(ret) - 1
	}
	return ret
}
