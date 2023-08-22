package binary_search

import (
	"math"
)

// https://leetcode.cn/problems/sqrtx/

// 方法一：数学公式
// 时间复杂度：O(logN)
// 空间复杂度：O(1)
func mySqrt0(x int) int {
	return int(math.Floor(math.Sqrt(float64(x))))
}

// 方法一：二分查找
// 时间复杂度：O(logN)
// 空间复杂度：O(1)
func mySqrt(x int) int {
	left, right := 0, x
	ans := 0
	for left <= right {
		mid := left + (right-left)/2
		if mid*mid <= x {
			left = mid + 1
			ans = mid
		} else {
			right = mid - 1
		}
	}
	return ans
}
