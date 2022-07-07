package main

import (
	"fmt"
	"math"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	var queue []int
	queue = append(queue, 1)
	queue = append(queue, 2)
	node := queue[0]
	queue = queue[1:]
	fmt.Println(node, queue)
	queue = append(queue, 3)
	queue = append(queue, 4)
	node = queue[0]
	queue = queue[1:]
	fmt.Println(node, queue)
}

func maxDepth1(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return int(math.Max(float64(maxDepth(root.Left)), float64(maxDepth(root.Right)))) + 1
}
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	var queue []*TreeNode
	node := root
	height := 1
	levelCount := 1
	queue = append(queue, root)
	for len(queue) > 0 {
		node = queue[0]
		queue = queue[1:]
		levelCount--
		if node.Left != nil {
			queue = append(queue, node.Left)
		}
		if node.Right != nil {
			queue = append(queue, node.Right)
		}
		if levelCount == 0 {
			height++
			levelCount = len(queue)
		}
	}
	return height
}
