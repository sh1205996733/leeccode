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
func levelOrder(root *TreeNode) (arr [][]int) {
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
			arr = append(arr, a)
			a = make([]int, 0)
		}
	}
	return arr
}

//迭代 内耗小
func levelOrder1(root *TreeNode) (arr [][]int) {
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
			arr = append(arr, a)
			a = make([]int, 0)
		}
	}
	return arr
}
