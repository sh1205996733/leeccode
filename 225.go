package main

import (
	"container/list"
	"fmt"
)

func main() {
	obj := Constructor()
	obj.Push(10)
	obj.Push(11)
	obj.Push(12)
	obj.Push(14)
	param_3 := obj.Top()
	param_2 := obj.Pop()
	param_4 := obj.Empty()
	fmt.Println(param_3, param_2, param_4)
}

//leetcode submit region begin(Prohibit modification and deletion)
type MyStack struct {
	*list.List
}

/** Initialize your data structure here. */
func Constructor() MyStack {
	l := new(list.List)
	return MyStack{l.Init()}
}

/** Push element x onto stack. */
func (this *MyStack) Push(x int) {
	this.PushBack(x)
}

/** Removes the element on top of the stack and returns that element. */
func (this *MyStack) Pop() int {
	return this.Remove(this.Back()).(int)
}

/** Get the top element. */
func (this *MyStack) Top() int {
	return this.Back().Value.(int)
}
func (this *MyStack) Empty() bool {
	return this.Len() == 0
}
