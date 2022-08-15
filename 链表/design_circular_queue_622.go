package linkedlist

//https://leetcode.cn/problems/design-circular-queue/
// 设计循环队列

type MyCircularQueue struct {
	data        []int
	front, tail int
}

func ConstructorMyCircularQueue(k int) MyCircularQueue {
	return MyCircularQueue{
		data: make([]int, k+1),
	}
}

func (this *MyCircularQueue) EnQueue(value int) bool {
	if this.IsFull() {
		return false
	}
	this.data[this.tail] = value
	this.tail = (this.tail + 1) % len(this.data)
	return true
}

func (this *MyCircularQueue) DeQueue() bool {
	if this.IsEmpty() {
		return false
	}
	this.data[this.front%len(this.data)] = -1
	this.front = (this.front + 1) % len(this.data)
	return true
}

func (this *MyCircularQueue) Front() int {
	if this.IsEmpty() {
		return -1
	}
	return this.data[this.front]
}

func (this *MyCircularQueue) Rear() int {
	if this.IsEmpty() {
		return -1
	}
	return this.data[(this.tail-1+len(this.data))%len(this.data)]
}

func (this *MyCircularQueue) IsEmpty() bool {
	return this.front == this.tail
}

func (this *MyCircularQueue) IsFull() bool {
	return (this.tail+1)%len(this.data) == this.front
}
