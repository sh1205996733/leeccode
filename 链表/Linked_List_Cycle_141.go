package linkedlist

// https://leetcode.cn/problems/linked-list-cycle/
// 环形链表

//  使用map存储已经遍历过的节点 每次遍历到一个节点时，判断该节点此前是否被访问过
// 时间复杂度：O(N)
// 空间复杂度：O(N)
func hasCycle0(head *ListNode) bool {
	if head == nil {
		return false
	}
	m := make(map[*ListNode]bool)
	m[head] = true

	for head.Next != nil {
		if ok, _ := m[head.Next]; ok {
			return true
		}
		m[head.Next] = true
		head = head.Next
	}
	return false
}

// 快慢指针 龟兔赛跑
// 时间复杂度：O(N)
// 空间复杂度：O(1)
func hasCycle(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}
	slow, fast := head, head.Next
	for fast != nil && fast.Next != nil {
		if fast == slow {
			return true
		}
		slow = slow.Next
		fast = fast.Next.Next
	}
	return false
}
