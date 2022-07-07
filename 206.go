package main

import "fmt"

func main() {
	//head := &ListNode{1,&ListNode{2,&ListNode{3,&ListNode{4,&ListNode{5,&ListNode{6,nil}}}}}}
	var head *ListNode
	head.string() //1_2_3_4_5_
	head = reverseList(head)
	head.string()
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func (node *ListNode) string() {
	for node != nil {
		fmt.Printf("%d_", node.Val)
		node = node.Next
	}
	fmt.Println()
}

//递归
func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	newHead := reverseList(head.Next) //此时newHead最后一个元素和head.next是同一个位置
	head.Next.Next = head             //newHead最后一个元素的Next指向head
	head.Next = nil                   //head的Next置空
	return newHead
}

//头插法(每次都往头部添加) 1_2_3_4_5_
func reverseList2(head *ListNode) *ListNode {
	var newHead *ListNode
	for head != nil {
		next := head.Next
		head.Next = newHead
		newHead = head
		head = next
	}
	return newHead
}
