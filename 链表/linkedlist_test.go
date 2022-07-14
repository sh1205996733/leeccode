package linkedlist

import "testing"

// 两数之和
func TestAddTwoNumbers(t *testing.T) {
	l1 := MakeLinkedList([]int{2, 4, 3})
	l2 := MakeLinkedList([]int{5, 6, 4})
	l3 := addTwoNumbers(l1, l2)
	PrintLinkedList(l3)
}

// 删除链表的倒数第 N 个结点
func TestRemoveNthFromEnd(t *testing.T) {
	l1 := MakeLinkedList([]int{1, 2, 3, 4, 5})
	l3 := removeNthFromEnd(l1, 2)
	PrintLinkedList(l3)
}

// 合并两个有序链表
func TestMergeTwoLists(t *testing.T) {
	l1 := MakeLinkedList([]int{3})
	l2 := MakeLinkedList([]int{2})
	l3 := mergeTwoLists(l1, l2)
	PrintLinkedList(l3)
}

// 排序链表
func TestSortList(t *testing.T) {
	l1 := MakeLinkedList([]int{4, 2, 1, 3})
	l3 := sortList(l1)
	PrintLinkedList(l3)
}
