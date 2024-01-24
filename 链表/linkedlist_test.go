package linkedlist

import (
	"fmt"
	binarytree "leetcode/二叉树"
	"leetcode/二叉树/printer"
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

// K 个一组翻转链表
func TestReverseKGroup(t *testing.T) {
	l1 := MakeLinkedList([]int{1, 2, 3, 4, 5, 6, 7, 8})
	l3 := reverseKGroup(l1, 3)
	PrintLinkedList(l3)
}

// 旋转链表
func TestRotateRight(t *testing.T) {
	l1 := MakeLinkedList([]int{0, 1, 2})
	l3 := rotateRight(l1, 4)
	PrintLinkedList(l3)
}

// 删除排序链表中的重复元素
func TestDeleteDuplicates(t *testing.T) {
	l1 := MakeLinkedList([]int{1, 1, 2, 3, 3})
	l3 := deleteDuplicates(l1)
	PrintLinkedList(l3)
}

// 删除排序链表中的重复元素II
func TestDeleteDuplicatesII(t *testing.T) {
	l1 := MakeLinkedList([]int{})
	l3 := deleteDuplicatesII(l1)
	PrintLinkedList(l3)
}

// 相交链表(两个链表的第一个公共节点)
func TestGetIntersectionNode(t *testing.T) {
	l1 := MakeLinkedList([]int{4, 1, 8, 4, 5})
	l2 := MakeLinkedList([]int{5, 0, 1, 8, 4, 5})
	l3 := getIntersectionNode(l1, l2)
	PrintLinkedList(l3)
}

// 分割链表
func TestPartition(t *testing.T) {
	l1 := MakeLinkedList([]int{4, 0, 1, 2})
	l3 := partition(l1, 3)
	PrintLinkedList(l3)
}

// 反转链表II
func TestReverseBetween(t *testing.T) {
	l1 := MakeLinkedList([]int{1})
	l3 := reverseBetween(l1, 1, 1)
	PrintLinkedList(l3)
}

// 有序链表转换二叉搜索树
func TestSortedListToBST(t *testing.T) {
	l1 := MakeLinkedList([]int{-10, -3, 0, 5, 9})
	l3 := sortedListToBST(l1)
	printer.Println(l3)
}

// 填充每个节点的下一个右侧节点指针
func TestConnect(t *testing.T) {
	l1 := binarytree.MakePerfectBinaryTree([]any{1, 2, 3, 4, 5, 6, 7})
	l3 := connect(l1)
	binarytree.PrintLinkedList2(l3)
}

// 填充每个节点的下一个右侧节点指针II
func TestConnectII(t *testing.T) {
	l1 := binarytree.MakePerfectBinaryTree([]any{1, 2, 3, 4, 5, 6, 7})
	l3 := connectII(l1)
	binarytree.PrintLinkedList2(l3)
}

// 复制带随机指针的链表
func TestCopyRandomList(t *testing.T) {
	//l1 := MakeLinkedList([]int{1, 2, 3, 4, 5, 6, 7})
	//l3 := copyRandomList(l1)
	//PrintLinkedList(l3)
}

// 环形链表
func TestHasCycle(t *testing.T) {
	l1 := MakeLinkedList([]int{3, 2, 0, -4})
	fmt.Println(hasCycle(l1))
}

// 环形链表II
func TestDetectCycle(t *testing.T) {
	l1 := MakeLinkedList([]int{3, 2, 0, -4})
	l3 := detectCycle(l1)
	PrintLinkedList(l3)
}

// 重排链表
func TestReorderList(t *testing.T) {
	l1 := MakeLinkedList([]int{1})
	reorderList(l1)
	PrintLinkedList(l1)
}

// 对链表进行插入排序
func TestInsertionSortList(t *testing.T) {
	l1 := MakeLinkedList([]int{2, 2, 2, 2})
	PrintLinkedList(insertionSortList(l1))
}

// 移除链表元素
func TestRemoveElements(t *testing.T) {
	l1 := MakeLinkedList([]int{})
	l3 := removeElements(l1, 0)
	PrintLinkedList(l3)
}

// 反转链表
func TestReverseList(t *testing.T) {
	l1 := MakeLinkedList([]int{1, 2, 3, 4, 5})
	PrintLinkedList(reverseList(l1))
}

// 删除链表中的节点
func TestDeleteNode(t *testing.T) {
	l1 := MakeLinkedList([]int{1})
	deleteNode(l1)
	PrintLinkedList(l1)
}

// 回文链表
func TestIsPalindrome(t *testing.T) {
	l1 := MakeLinkedList([]int{1, 0, 0, 1})
	fmt.Println(isPalindrome(l1))
}

// 奇偶链表
func TestOddEvenList(t *testing.T) {
	l1 := MakeLinkedList([]int{1})
	PrintLinkedList(oddEvenList(l1))
}

// 分隔链表
func TestSplitListToParts(t *testing.T) {

	l1 := MakeLinkedList([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})
	l3 := splitListToParts(l1, 4)
	for _, v := range l3 {
		PrintLinkedList(v)
	}
}

// 链表随机节点
func TestSolution(t *testing.T) {
	l1 := MakeLinkedList([]int{1})
	obj := ConstructorSolution(l1)
	fmt.Println(obj.GetRandom())
	fmt.Println(obj.GetRandom())
	fmt.Println(obj.GetRandom())
	fmt.Println(obj.GetRandom())
	fmt.Println(obj.GetRandom())
}

// 链表中的下一个更大节点
func TestNextLargerNodes(t *testing.T) {
	l1 := MakeLinkedList([]int{9, 7, 6, 7, 6, 9})
	fmt.Println(nextLargerNodes(l1))
}

// 两数相加II
func TestAddTwoNumbersII(t *testing.T) {
	l1 := MakeLinkedList([]int{7, 2, 4, 3})
	l2 := MakeLinkedList([]int{5, 6, 4})
	PrintLinkedList(addTwoNumbersII(l1, l2))
}

// 从链表中删去总和值为零的连续节点
func TestRemoveZeroSumSublists(t *testing.T) {
	l1 := MakeLinkedList([]int{1, 2, -103, 4, -4, -3, 2, 1})
	PrintLinkedList(removeZeroSumSublists(l1))
}

// 链表组件
func TestNumComponents(t *testing.T) {
	l1 := MakeLinkedList([]int{0, 1, 2, 4, 3})
	fmt.Println(numComponents(l1, []int{0, 1, 3}))
}

// 单链表加1
func TestPlusOne(t *testing.T) {
	l1 := MakeLinkedList([]int{9, 9, 9})
	PrintLinkedList(plusOne(l1))
}

// 二进制链表转整数
func TestGetDecimalValue(t *testing.T) {
	l1 := MakeLinkedList([]int{1, 0, 1, 1, 0})
	fmt.Println(getDecimalValue(l1))
}

func Test(t *testing.T) {
	MakeLinkedList([]int{1, 2, 3, 4, 5})
}
