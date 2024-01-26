package binary_search

// 搜索插入位置
// https://leetcode.cn/problems/search-insert-position/

// 方法一：二分查找
// 时间复杂度：O(logN)
// 空间复杂度：O(1)
func searchInsert(nums []int, target int) int {
	//begin, end := 0, len(nums)
	//for begin < end {
	//	mid := (begin + end) >> 1
	//	if nums[mid] > target {
	//		end = mid
	//	} else if nums[mid] == target {
	//		return mid
	//	} else {
	//		begin = mid + 1
	//	}
	//}
	//return begin

	begin, end := 0, len(nums)-1
	for begin <= end {
		mid := (begin + end) >> 1
		if target <= nums[mid] {
			end = mid - 1
		} else {
			begin = mid + 1
		}
	}
	return begin
}
