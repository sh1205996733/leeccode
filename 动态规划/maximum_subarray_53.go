package dp

import (
	"fmt"
	"math"
)

// https://leetcode.cn/problems/maximum-subarray/

// 方法一: 递归
// 时间复杂度：O(n)
// 空间复杂度：O(n)
// 分治
func maxSubArray(nums []int) int {
	if nums == nil || len(nums) == 0 {
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

// 动态规划
func maxSubArray0(nums []int) interface{} {
	if nums == nil || len(nums) == 0 {
		return 0
	}
	dp := nums[0]
	begin, end := 0, 0
	max := dp
	for i := 1; i < len(nums); i++ {
		if dp > 0 {
			dp = dp + nums[i]
		} else {
			dp = nums[i]
			begin = i
		}
		if dp > max {
			max = dp
			end = i
		}
	}
	fmt.Println(begin, end, nums[begin:end+1])
	return max
}
