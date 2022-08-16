package linkedlist

// https://leetcode.cn/problems/rotate-list/
// 旋转链表

// 方法一: 迭代 先遍历链表 求出size,并将首尾相连组成一个环 然后head 向右移动 size - k % size -1 个位置 最后拼接成一个新的head
// 时间复杂度为 O(n)
// 空间复杂度：O(1)
func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil {
		return head
	}
	cur := head
	size := 0
	for cur != nil {
		size++
		if cur.Next == nil {
			cur.Next = head
			break
		}
		cur = cur.Next
	}
	move := size - k%size
	cur = head
	for i := 1; i < move; i++ {
		cur = cur.Next
	}
	//设置新的head
	head = cur.Next
	cur.Next = nil
	return head
}
