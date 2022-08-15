package linkedlist

// https://leetcode.cn/problems/linked-list-components/
// 链表组件

// 方法一：把nums放到map，在遍历head
// 时间复杂度：O(N)
// 空间复杂度：O(N)
func numComponents(head *ListNode, nums []int) int {
	m := map[int]bool{}
	for _, num := range nums {
		m[num] = true
	}
	total := 0
	for node := head; node != nil; node = node.Next {
		if m[node.Val] {
			cur := node.Next
			for ; cur != nil && m[cur.Val]; cur = cur.Next {

			}
			total++
			node = cur
		}
		if node == nil {
			break
		}
	}
	return total
}
