package linkedlist

// https://leetcode.cn/problems/partition-list/
// 分隔链表

// 方法一: 迭代 head有可能会改变 所以使用一个虚拟的头节点，然后记录插入点位置tail，找到符合的节点 再插入到tail之后，然后更新一下tail
// 时间复杂度为 O(n)
// 空间复杂度：O(1)
func partition0(head *ListNode, x int) *ListNode {
	dummyHead := &ListNode{-101, head}
	head = dummyHead
	for cur := dummyHead; cur != nil && cur.Next != nil; cur = cur.Next {
		for next := cur.Next; next != nil && next.Val < x; next = cur.Next {
			cur.Next = next.Next
			if cur == head {
				cur = next
			}
			next.Next = head.Next
			head.Next = next
			head = next
		}

	}
	return dummyHead.Next
}

// 直观来说我们只需维护两个链表 small 和 large 即可，small 链表按顺序存储所有小于 xx 的节点，large 链表按顺序存储所有大于等于 xx 的节点。
// 遍历完原链表后，我们只要将 small 链表尾节点指向 large 链表的头节点即能完成对链表的分隔。

func partition(head *ListNode, x int) *ListNode {
	smallHead := &ListNode{}
	largeHead := &ListNode{}
	smallTail, largeTail := smallHead, largeHead
	for ; head != nil; head = head.Next {
		if head.Val < x {
			smallTail.Next = head
			smallTail = smallTail.Next
		} else {
			largeTail.Next = head
			largeTail = largeTail.Next
		}
	}
	largeTail.Next = nil
	smallTail.Next = largeHead.Next
	return smallHead.Next
}
