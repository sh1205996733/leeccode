package binary_search

// 寻找旋转排序数组中的最小值
// https://leetcode.cn/problems/find-minimum-in-rotated-sorted-array/description/
func findMin(nums []int) int {
	le, ri := 0, len(nums)
	for le < ri {
		mid := le + (ri-le)/2
		if nums[mid] <= nums[len(nums)-1] {
			ri = mid
		} else {
			le = mid + 1
		}
	}
	return nums[le]
}
