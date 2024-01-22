package binarytree

// 对称二叉树
// https://leetcode.cn/problems/symmetric-tree/description/

// 方法一: 递归
func isSymmetric0(root *TreeNode) bool {
	if root == nil {
		return false
	}
	return symmetric(root.Left, root.Right)
}

func symmetric(left *TreeNode, right *TreeNode) bool {
	if left == nil && right == nil {
		return true
	} else if left == nil || right == nil || left.Val != right.Val {
		return false
	}
	return symmetric(left.Left, right.Right) && symmetric(left.Right, right.Left)
}

// 方法二: 迭代
func isSymmetric1(root *TreeNode) bool {
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

// 方法三:队列
func isSymmetric(root *TreeNode) bool {
	u, v := root, root
	var queue []*TreeNode
	queue = append(queue, root, root)
	for len(queue) > 0 {
		u, v = queue[0], queue[1]
		queue = queue[2:]
		if u == nil && v == nil {
			continue
		}
		if u == nil || v == nil || u.Val != v.Val {
			return false
		}
		queue = append(queue, u.Left, v.Right, u.Right, v.Left)
	}
	return true
}
