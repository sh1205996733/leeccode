package binary_search

// 搜索旋转排序数组
// https://leetcode.cn/problems/search-in-rotated-sorted-array/description/
func search2(nums []int, target int) int {
	le, ri := 0, len(nums)-1
	for le <= ri {
		mid := le + (ri-le)/2
		if nums[mid] == target {
			return mid
		}
		if nums[mid] < nums[ri] { //后半部分有序
			if nums[mid] < target && target <= nums[ri] {
				le = mid + 1
			} else {
				ri = mid - 1
			}
		} else {
			if nums[mid] > target && target >= nums[le] {
				ri = mid - 1
			} else {
				le = mid + 1
			}
		}
	}
	return -1
}
