package monotonicstack

import "fmt"

// https://leetcode.cn/problems/trapping-rain-water/
// 接雨水

// 方法一: 暴力法 按列求 找当前柱子左右最大值
// 求每一列的水，我们只需要关注当前列，以及左边最高的墙，右边最高的墙就够了。 装水的多少，当然根据木桶效应，我们只需要看左边最高的墙和右边最高的墙中较矮的一个就够了。
// 时间复杂度O(n^2)
// 空间复杂度O(1)
func trap0(height []int) int {
	ans := 0
	for i, h := range height {
		leftMax := 0
		for l := i - 1; l >= 0; l-- { //左边最大值
			if height[l] > leftMax {
				leftMax = height[l]
			}
		}
		rightMax := 0
		for r := i + 1; r < len(height); r++ { //右边最大值
			if height[r] > rightMax {
				rightMax = height[r]
			}
		}
		min := Min(leftMax, rightMax)
		if min > h {
			ans += min - h
		}
	}
	return ans
}

// 方法二: 动态规划 优化方案一(存储最大高度)
// 时间复杂度O(n)
// 空间复杂度O(n)
func trap(height []int) int {
	ans := 0
	n := len(height)
	leftMax, rightMax := make([]int, n), make([]int, n)
	leftMax[0], rightMax[n-1] = height[0], height[n-1]
	for i := 1; i < len(height); i++ {
		leftMax[i] = max(leftMax[i-1], height[i])
	}
	for i := len(height) - 2; i >= 0; i-- {
		rightMax[i] = max(rightMax[i+1], height[i])
	}
	for i, h := range height {
		ans += Min(leftMax[i], rightMax[i]) - h
	}
	return ans
}

// 方法四: 单调栈 找第一个比他大的数
// 时间复杂度O(n)
// 空间复杂度O(n)
func trap2(height []int) int {
	ans := 0
	var stack []int
	for i := range height {
		for len(stack) > 0 && height[stack[len(stack)-1]] < height[i] {
			h := height[stack[len(stack)-1]]
			stack = stack[:len(stack)-1]
			if len(stack) == 0 {
				break
			}
			min := Min(height[stack[len(stack)-1]], height[i])
			d := i - stack[len(stack)-1] - 1
			ans += d * (min - h)
			fmt.Printf("%d %d\n", d, h)
		}
		stack = append(stack, i)
	}
	return ans
}
func Min(a, b int) int {
	if b > a {
		return a
	}
	return b
}
func Max(a, b int) int {
	if b > a {
		return b
	}
	return a
}
