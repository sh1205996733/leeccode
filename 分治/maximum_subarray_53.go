package divide_conquer

import (
	"math"
)

// https://leetcode.cn/problems/maximum-subarray/

// 方法一: 分治
// 时间复杂度：O(n)
// 空间复杂度：O(n)
func maxSubArray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	return maxArray(nums, 0, len(nums))
}

func maxArray(nums []int, begin int, end int) int {
	if end-begin < 2 {
		return nums[begin]
	}
	mid := (end + begin) >> 1
	leftMax := nums[mid-1]
	leftSum := leftMax
	for i := mid - 2; i >= begin; i-- {
		leftSum += nums[i]
		if leftSum > leftMax {
			leftMax = leftSum
		}
	}
	rightMax := nums[mid]
	rightSum := rightMax
	for i := mid + 1; i < end; i++ {
		rightSum += nums[i]
		if rightSum > rightMax {
			rightMax = rightSum
		}
	}
	return int(math.Max(float64(leftMax+rightMax), math.Max(float64(maxArray(nums, begin, mid)), float64(maxArray(nums, mid, end)))))
}
