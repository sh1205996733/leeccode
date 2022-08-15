package linkedlist

//https://leetcode.cn/problems/design-linked-list/
// 设计链表

type MyLinkedList struct {
	dammyHead *ListNode
	tail      *ListNode
	size      int
}

func ConstructorMyLinkedList() MyLinkedList {
	return MyLinkedList{dammyHead: &ListNode{}, size: 0}
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
		this.dammyHead.Next = &ListNode{val, this.dammyHead.Next}
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
		this.dammyHead.Next = this.dammyHead.Next.Next
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
	p := this.dammyHead
	for i := 0; i <= index; i++ {
		p = p.Next
	}
	return p
}
