package binary_search

// https://leetcode.cn/problems/find-peak-element/
// 因为左右都是负无穷大且不存在相邻的值是一样的 所以直接遍历（找到最大值）
// 方法一：暴力遍历 （找最大值）
// 时间复杂度：O(N)
// 空间复杂度：O(1)
func findPeakElement0(nums []int) int {
	idx := 0
	for i, v := range nums {
		if v > nums[idx] {
			idx = i
		}
	}
	return -1
}

// 方法二：二分查找 只要找到上升、下降的必然符合
// 时间复杂度：O(logN)
// 空间复杂度：O(1)
func findPeakElement(nums []int) int {
	le, ri := 0, len(nums)-1
	for le < ri {
		mid := le + (ri-le)/2
		if nums[mid] > nums[mid+1] {
			ri = mid
		} else {
			le = mid + 1
		}
	}
	return le
}
