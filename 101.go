package main

import "fmt"

func main() {
	a := []int{1, 3, 4}
	fmt.Println(a[1:])
	fmt.Println(a[:1])
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSymmetric(root *TreeNode) bool {
	var queue []*TreeNode //模拟队列
	var nums []*TreeNode
	node := root
	queue = append(queue, node)
	levelCount := 1
	for len(queue) > 0 {
		node = queue[0] //头
		queue = queue[1:]
		levelCount--
		if node.Left != nil {
			queue = append(queue, node.Left)
			nums = append(nums, node.Left)
		} else {
			nums = append(nums, nil)
		}
		if node.Right != nil {
			queue = append(queue, node.Right)
			nums = append(nums, node.Right)
		} else {
			nums = append(nums, nil)
		}
		if levelCount == 0 {
			levelCount = len(queue)
			length := len(nums)
			mid := length / 2
			for i := 0; i < mid; i++ {
				if nums[i] == nil && nums[length-1-i] == nil {
					continue
				}
				if nums[i] == nil && nums[length-1-i] != nil || nums[i] != nil && nums[length-1-i] == nil || nums[i].Val != nums[length-1-i].Val {
					return false
				}
			}
			nums = make([]*TreeNode, 0)
		}
	}
	return true
}

//递归
func isSymmetric1(root *TreeNode) bool {
	return symmetric(root.Left, root.Right)
}
func symmetric(left *TreeNode, right *TreeNode) bool {
	if left == nil && right == nil {
		return true
	} else if left != nil && right != nil && left.Val == right.Val {
		return symmetric(left.Left, right.Right) && symmetric(left.Right, right.Left)
	}
	return false
}

func isSymmetric(root *TreeNode) bool {
	u, v := root, root
	q := []*TreeNode{}
	q = append(q, u)
	q = append(q, v)
	for len(q) > 0 {
		u, v = q[0], q[1]
		q = q[2:]
		if u == nil && v == nil {
			continue
		}
		if u == nil || v == nil {
			return false
		}
		if u.Val != v.Val {
			return false
		}
		q = append(q, u.Left)
		q = append(q, v.Right)

		q = append(q, u.Right)
		q = append(q, v.Left)
	}
	return true
}
