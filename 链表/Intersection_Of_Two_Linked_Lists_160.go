package linkedlist

// https://leetcode.cn/problems/intersection-of-two-linked-lists/
// 相交链表(两个链表的第一个公共节点)

//方法一 :迭代 先求出长度 在将长的移动max-min 在同时开始遍历两个链表
// 时间复杂度为 O(n)
// 空间复杂度：O(1)
func getIntersectionNode1(headA, headB *ListNode) *ListNode {
	lA, lB := getLen(headA), getLen(headB)
	max, min := headA, headB
	index := lA - lB
	if lA < lB {
		max, min = min, max
		index = lB - lA
	}
	for max != nil && min != nil {
		for index > 0 {
			max = max.Next
			index--
		}
		if max == min {
			return max
		}
		max = max.Next
		min = min.Next
	}
	return nil
}

// 方法二: 使用map
// 时间复杂度为 O(n)
// 空间复杂度：O(n)
func getIntersectionNode2(headA, headB *ListNode) *ListNode {
	m := map[*ListNode]bool{}
	head := headA
	for head != nil {
		m[head] = true
		head = head.Next
	}
	head = headB
	for head != nil {
		if m[head] {
			return head
		}
		head = head.Next
	}
	return nil
}

func getLen(a *ListNode) int {
	next := a
	count := 0
	for next != nil {
		count++
		next = next.Next
	}
	return count
}

//方法一 :迭代 首位相连
// 时间复杂度为 O(n)
// 空间复杂度：O(1)
// 4,1,8,4,5,5,0,1,8,4,5
// 5,0,1,8,4,5,4,1,8,4,5
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	pa := headA
	pb := headB
	for pa != pb {
		if pa == nil {
			pa = headB
		} else {
			pa = pa.Next
		}
		if pb == nil {
			pb = headA
		} else {
			pb = pb.Next
		}
	}
	return pa
}
