package main

import (
	"fmt"
	"math"
)

func main() {
	nums := []int{-10, -3, 0, 5, 9}
	fmt.Printf("%+v", sortedArrayToBST(nums))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//-10,-3,0,5,9
func sortedArrayToBST(nums []int) *TreeNode {
	return buildTree(nums, 0, len(nums)-1)
}

func buildTree(nums []int, begin int, end int) *TreeNode {
	if begin > end { //begin > end
		return nil
	}
	mid := int(math.Ceil(float64(begin + (end-begin)>>1)))
	if mid == begin {
		mid = end
	}
	return &TreeNode{nums[mid], buildTree(nums, begin, mid-1), buildTree(nums, mid+1, end)}
}
