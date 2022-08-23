package queue

// https://leetcode.cn/problems/implement-stack-using-queues/
// 用队列实现栈

// 方法一：使用双端队列
// 时间复杂度：O(N)
// 空间复杂度：O(N)

//type MyStack struct {
//	queue *list.List
//}
//
//func Constructor() MyStack {
//	return MyStack{
//		list.New(),
//	}
//}
//
//func (this *MyStack) Push(x int) {
//	this.queue.PushBack(x)
//}
//
//func (this *MyStack) Pop() int {
//	back := this.queue.Back()
//	this.queue.Remove(back)
//	return back.Value.(int)
//}
//
//func (this *MyStack) Top() int {
//	return this.queue.Back().Value.(int)
//}
//
//func (this *MyStack) Empty() bool {
//	return this.queue.Len() == 0
//}

// 方法二: 使用双队列 queue2辅助存储 queue1主存储
// 时间复杂度：O(N)
// 空间复杂度：O(N)

type MyStack struct {
	queue1 []int
	queue2 []int
}

func Constructor() MyStack {
	return MyStack{
		[]int{},
		[]int{},
	}
}

func (this *MyStack) Push(x int) {
	this.queue2 = append(this.queue2, x)
	for len(this.queue1) > 0 {
		this.queue2 = append(this.queue2, this.queue1[0])
		this.queue1 = this.queue1[1:]
	}
	this.queue1, this.queue2 = this.queue2, this.queue1
}

func (this *MyStack) Pop() int {
	val := this.queue1[0]
	this.queue1 = this.queue1[1:]
	return val
}

func (this *MyStack) Top() int {
	return this.queue1[0]
}

func (this *MyStack) Empty() bool {
	return len(this.queue1) == 0
}
