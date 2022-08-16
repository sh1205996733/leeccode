package linkedlist

// https://leetcode.cn/problems/remove-zero-sum-consecutive-nodes-from-linked-list/
// 从链表中删去总和值为零的连续节点

// 方法一：暴力法
// 时间复杂度：O(N^2)
// 空间复杂度：O(1)
func removeZeroSumSublists0(head *ListNode) *ListNode {
	dummyHead := &ListNode{Next: head}
	sum := 0
	node := head
	for cur := head; cur != nil; cur = cur.Next {
		prev := dummyHead
		sum += cur.Val
		if sum != 0 {
			node = prev.Next
			sum1 := sum - cur.Val
			for node != cur {
				if sum1+cur.Val != 0 {
					sum1 -= node.Val
					prev = node
					node = node.Next
				} else {
					prev.Next = cur.Next
					break
				}
			}
		} else {
			prev.Next = cur.Next
		}
	}
	return dummyHead.Next
}

// 方法二：两个map
// 时间复杂度：O(N)
// 空间复杂度：O(N)
func removeZeroSumSublists(head *ListNode) *ListNode {
	dummyHead := &ListNode{Next: head}
	m1 := map[int]*ListNode{}
	sum := 0
	for p := dummyHead; p != nil; p = p.Next {
		sum += p.Val
		m1[sum] = p
	}
	sum2 := 0
	for p := dummyHead; p != nil; p = p.Next {
		sum2 += p.Val
		p.Next = m1[sum2].Next
	}
	return dummyHead.Next
}
