package linkedlist

// https://leetcode.cn/problems/linked-list-cycle-ii/
// https://leetcode.cn/problems/linked-list-cycle-lcci/
// 环形链表II

//  使用map存储已经遍历过的节点 每次遍历到一个节点时，判断该节点此前是否被访问过
// 时间复杂度：O(N)
// 空间复杂度：O(N)
func detectCycle0(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return nil
	}
	m := make(map[*ListNode]bool)
	for head != nil {
		if ok, _ := m[head]; ok {
			return head
		}
		m[head] = true
		head = head.Next
	}
	return nil
}

// 快慢指针 龟兔赛跑 当发现 slow 与 相遇时，我们再额外使用一个指针 ptr。起始，它指向链表头部；随后，它和 slow 每次向后移动一个位置。最终，它们会在入环点相遇
// 时间复杂度：O(N)
// 空间复杂度：O(1)
func detectCycle(head *ListNode) *ListNode {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if fast == slow {
			p := head
			for p != slow {
				p = p.Next
				slow = slow.Next
			}
			return p
		}
	}
	return nil
}
