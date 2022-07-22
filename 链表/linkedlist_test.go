package linkedlist

import (
	"testing"
)

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
	l1 := MakeLinkedList([]int{4, 2, 1, 3, 5})
	l3 := sortList(l1)
	PrintLinkedList(l3)
}

// 合并K个升序链表
func TestMergeKLists(t *testing.T) {
	data := [][]int{{}, {1}}
	var list []*ListNode
	for _, d := range data {
		list = append(list, MakeLinkedList(d))
	}
	l3 := mergeKLists(list)
	PrintLinkedList(l3)
}

// 两两交换链表中的节点
func TestWwapPairs(t *testing.T) {
	l1 := MakeLinkedList([]int{1, 2, 3, 4, 5})
	l3 := swapPairs(l1)
	PrintLinkedList(l3)
}

//  K 个一组翻转链表
func TestReverseKGroup(t *testing.T) {
	l1 := MakeLinkedList([]int{1, 2, 3, 4, 5, 6, 7, 8})
	l3 := reverseKGroup(l1, 3)
	PrintLinkedList(l3)
}

//  旋转链表
func TestRotateRight(t *testing.T) {
	l1 := MakeLinkedList([]int{0, 1, 2})
	l3 := rotateRight(l1, 4)
	PrintLinkedList(l3)
}

//  删除排序链表中的重复元素
func TestDeleteDuplicates(t *testing.T) {
	l1 := MakeLinkedList([]int{1, 1, 2, 3, 3})
	l3 := deleteDuplicates(l1)
	PrintLinkedList(l3)
}

//  删除排序链表中的重复元素II
func TestDeleteDuplicatesII(t *testing.T) {
	l1 := MakeLinkedList([]int{2, 3, 3})
	l3 := deleteDuplicatesII(l1)
	PrintLinkedList(l3)
}

//  删除排序链表中的重复元素II
func TestGetIntersectionNode(t *testing.T) {
	l1 := MakeLinkedList([]int{4, 1, 8, 4, 5})
	l2 := MakeLinkedList([]int{5, 0, 1, 8, 4, 5})
	l3 := getIntersectionNode(l1, l2)
	PrintLinkedList(l3)
}
