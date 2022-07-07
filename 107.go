package main

import (
	"container/list"
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	var a []int
	fmt.Printf("%p\n", a)
	a = append(a, 1)
	a = a[:0]
	fmt.Printf("%p\n", a)
	fmt.Println(a)
}

//迭代 内耗大
func levelOrderBottom(root *TreeNode) (ret [][]int) {
	if root == nil {
		return [][]int{}
	}
	queue := list.New()
	node := root
	queue.PushBack(node)
	levelCount := 1
	var a []int
	for queue.Len() > 0 {
		node = queue.Remove(queue.Front()).(*TreeNode)
		a = append(a, node.Val)
		levelCount--
		if node.Left != nil {
			queue.PushBack(node.Left)
		}
		if node.Right != nil {
			queue.PushBack(node.Right)
		}
		if levelCount == 0 {
			levelCount = queue.Len()
			t := [][]int{a}
			ret = append(t, ret...)
			a = make([]int, 0)
		}
	}
	return ret
}

//迭代 内耗小
func levelOrderBottom(root *TreeNode) (ret [][]int) {
	if root == nil {
		return [][]int{}
	}
	var queue []*TreeNode //模拟队列
	node := root
	queue = append(queue, node)
	levelCount := 1
	var a []int
	for len(queue) > 0 {
		node = queue[0]
		queue = queue[1:]
		a = append(a, node.Val)
		levelCount--
		if node.Left != nil {
			queue = append(queue, node.Left)
		}
		if node.Right != nil {
			queue = append(queue, node.Right)
		}
		if levelCount == 0 {
			levelCount = len(queue)
			ret = append(ret, a)
			a = make([]int, 0)
		}
	}
	start, end := 0, len(ret)-1
	for start < end {
		ret[start], ret[end] = ret[end], ret[start]
		start++
		end--
	}
	return ret
}
