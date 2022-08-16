package linkedlist

//https://leetcode.cn/problems/design-linked-list/
// 设计链表

type MyLinkedList struct {
	dummyHead *ListNode
	tail      *ListNode
	size      int
}

func ConstructorMyLinkedList() MyLinkedList {
	return MyLinkedList{dummyHead: &ListNode{}, size: 0}
}

func (this *MyLinkedList) Get(index int) int {
	node := this.node(index)
	if node == nil {
		return -1
	}
	return node.Val
}

func (this *MyLinkedList) AddAtHead(val int) {
	this.AddAtIndex(0, val)
}

func (this *MyLinkedList) AddAtTail(val int) {
	this.AddAtIndex(this.size, val)
}

func (this *MyLinkedList) AddAtIndex(index int, val int) {
	if index > this.size {
		return
	}
	if index <= 0 {
		this.dummyHead.Next = &ListNode{val, this.dummyHead.Next}
	} else {
		prev := this.node(index - 1)
		prev.Next = &ListNode{val, prev.Next}
	}
	this.size++
}

func (this *MyLinkedList) DeleteAtIndex(index int) {
	if index >= this.size {
		return
	}
	if index <= 0 {
		this.dummyHead.Next = this.dummyHead.Next.Next
	} else {
		prev := this.node(index - 1)
		prev.Next = prev.Next.Next
	}
	this.size--
}

func (this *MyLinkedList) node(index int) *ListNode {
	if index < 0 || index >= this.size {
		return nil
	}
	p := this.dummyHead
	for i := 0; i <= index; i++ {
		p = p.Next
	}
	return p
}
