package linkedlist

// https://leetcode.cn/problems/maximum-twin-sum-of-a-linked-list/
// 链表最大孪生和

// 放到数组中
// 时间复杂度：O(N)
// 空间复杂度：O(n)
func pairSum(head *ListNode) int {
	var nodes []*ListNode
	for node := head; node != nil; node = node.Next {
		nodes = append(nodes, node)
	}
	max, size := 0, len(nodes)
	for i := 0; i < size; i++ {
		t := nodes[i].Val + nodes[size-1-i].Val
		if t > max {
			max = t
		}
	}
	return max
}

// 快慢指针 找到中间元素，然后将后半部分反转
// 时间复杂度：O(N)
// 空间复杂度：O(1)
func pairSum1(head *ListNode) int {
	var nodes []*ListNode
	for node := head; node != nil; node = node.Next {
		nodes = append(nodes, node)
	}
	max, size := 0, len(nodes)
	for i := 0; i < size; i++ {
		t := nodes[i].Val + nodes[size-1-i].Val
		if t > max {
			max = t
		}
	}
	return max
}
