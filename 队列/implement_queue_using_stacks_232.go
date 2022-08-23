package queue

// https://leetcode.cn/problems/implement-queue-using-stacks/
// 用栈实现队列

// 方法一：双栈
// 时间复杂度：O(1)
// 空间复杂度：O(N)

type MyQueue struct {
	inStack, outStack []int
}

func ConstructorQueue() MyQueue {
	return MyQueue{
		[]int{},
		[]int{},
	}
}

func (p *MyQueue) Push(x int) {
	p.inStack = append(p.inStack, x)
}

func (p *MyQueue) Pop() int {
	if len(p.outStack) == 0 {
		for len(p.inStack) > 0 {
			p.outStack = append(p.outStack, p.inStack[len(p.inStack)-1])
			p.inStack = p.inStack[:len(p.inStack)-1]
		}
	}
	val := p.outStack[len(p.outStack)-1]
	p.outStack = p.outStack[:len(p.outStack)-1]
	return val
}

func (p *MyQueue) Peek() int {
	if len(p.outStack) == 0 {
		for len(p.inStack) > 0 {
			p.outStack = append(p.outStack, p.inStack[len(p.inStack)-1])
			p.inStack = p.inStack[:len(p.inStack)-1]
		}
	}
	return p.outStack[len(p.outStack)-1]
}

func (p *MyQueue) Empty() bool {
	return len(p.inStack) == 0 && len(p.outStack) == 0
}
