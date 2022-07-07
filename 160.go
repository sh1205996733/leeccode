/**
 * @Author:shangyc
 * @Date: 2022/5/12 0012
 */

package main

import "fmt"

func main() {
	intersectionNode := &ListNode{8, &ListNode{4, &ListNode{5, nil}}}
	headA := &ListNode{4, &ListNode{1, intersectionNode}}
	headB := &ListNode{5, &ListNode{6, &ListNode{1, intersectionNode}}}
	fmt.Println(getIntersectionNode(headA, headB))
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func getIntersectionNode1(headA, headB *ListNode) *ListNode {
	lA, lB := getLen(headA), getLen(headB)
	max, min := headA, headB
	index := lA - lB
	if lA < lB {
		max, min = min, max
		index = lB - lA
	}
	for max != nil && min != nil {
		for index > 0 {
			max = max.Next
			index--
		}
		if max == min {
			return max
		}
		max = max.Next
		min = min.Next
	}
	return nil
}
func getIntersectionNode2(headA, headB *ListNode) *ListNode {
	m := map[*ListNode]bool{}
	head := headA
	for head != nil {
		m[head] = true
		head = head.Next
	}
	head = headB
	for head != nil {
		if m[head] {
			return head
		}
		head = head.Next
	}
	return nil
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}
	pa, pb := headA, headB
	for pa != pb {
		if pa == nil {
			pa = headB
		} else {
			pa = pa.Next
		}
		if pb == nil {
			pb = headA
		} else {
			pb = pb.Next
		}
	}
	return pa
}

func getLen(a *ListNode) int {
	next := a
	count := 0
	for next != nil {
		count++
		next = next.Next
	}
	return count
}
