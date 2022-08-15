package linkedlist

//https://leetcode.cn/problems/design-circular-deque/
// 设计循环双端队列

type MyCircularDeque struct {
	element     []int
	front, tail int
}

func ConstructorMyCircularDeque(k int) MyCircularDeque {
	return MyCircularDeque{
		element: make([]int, k+1),
	}
}

func (this *MyCircularDeque) InsertFront(value int) bool {
	if this.IsFull() {
		return false
	}
	this.front = (this.front - 1 + len(this.element)) % len(this.element)
	this.element[this.front] = value
	return true
}

func (this *MyCircularDeque) InsertLast(value int) bool {
	if this.IsFull() {
		return false
	}
	this.element[this.tail] = value
	this.tail = (this.tail + 1) % len(this.element)
	return true
}

func (this *MyCircularDeque) DeleteFront() bool {
	if this.IsEmpty() {
		return false
	}
	this.element[this.front%len(this.element)] = -1
	this.front = (this.front + 1) % len(this.element)
	return true
}

func (this *MyCircularDeque) DeleteLast() bool {
	if this.IsEmpty() {
		return false
	}
	this.element[this.tail] = -1
	this.tail = (this.tail - 1 + len(this.element)) % len(this.element)
	return true
}

func (this *MyCircularDeque) GetFront() int {
	if this.IsEmpty() {
		return -1
	}
	return this.element[this.front]
}

func (this *MyCircularDeque) GetRear() int {
	if this.IsEmpty() {
		return -1
	}
	return this.element[(this.tail-1+len(this.element))%len(this.element)]
}

func (this *MyCircularDeque) IsEmpty() bool {
	return this.front == this.tail
}

func (this *MyCircularDeque) IsFull() bool {
	return (this.tail+1)%len(this.element) == this.front
}
