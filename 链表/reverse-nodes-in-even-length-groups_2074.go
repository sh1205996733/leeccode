package linkedlist

// https://leetcode.cn/problems/reverse-nodes-in-even-length-groups/
// 反转偶数长度组的节点

// 放到数组中
// 时间复杂度：O(N)
// 空间复杂度：O(1)
func reverseEvenLengthGroups(head *ListNode) *ListNode {
	var nodes []*ListNode
	for node, size := head, 1; node != nil; node = node.Next {
		nodes = append(nodes, node)
		if len(nodes) == size || node.Next == nil { // 统计到 size 个节点，或到达链表末尾
			if n := len(nodes); n%2 == 0 { // 有偶数个节点
				for i := 0; i < n/2; i++ {
					nodes[i].Val, nodes[n-1-i].Val = nodes[n-1-i].Val, nodes[i].Val // 直接交换元素值
				}
			}
			nodes = nil
			size++
		}
	}
	return head
}
