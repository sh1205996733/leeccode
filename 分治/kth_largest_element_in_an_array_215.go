package divide_conquer

// 数组中的第K个最大元素
// https://leetcode.cn/problems/kth-largest-element-in-an-array/description/

// 方法一：快速排序（分治）
// 时间复杂度：O(N)
// 空间复杂度：O(N)
func findKthLargest(nums []int, k int) int {
	quickKth(nums, 0, len(nums)-1, k)
	return 0
}
func quickKth(nums []int, le int, ri int, k int) {
	if le == ri {
		return
	}
	mid := (le + ri) >> 1
	quickKth(nums, le, mid, k)
	quickKth(nums, mid+1, ri, k)
	if nums[le] < nums[ri] {
		nums[le], nums[ri] = nums[ri], nums[le]
	}
}
