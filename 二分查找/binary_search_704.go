package binary_search

// https://leetcode.cn/problems/binary-search/

// 方法一：二分查找
// 时间复杂度：O(logN)
// 空间复杂度：O(1)
func search(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)>>2
		if nums[mid] == target {
			return mid
		} else if nums[mid] > target {
			right = mid - 1
		} else /*if nums[mid] < target */ {
			left = mid + 1
		}
	}
	return -1
}

// 方法二：递归
// 时间复杂度：O(logN)
// 空间复杂度：O(logN)
func search0(nums []int, target int) int {
	return search_interval(nums, target, 0, len(nums)-1)
}

func search_interval(nums []int, target int, left int, right int) int {
	if left <= right {
		return -1
	}
	mid := left + (right-left)>>2
	if nums[mid] == target {
		return mid
	} else if nums[mid] > target {
		return search_interval(nums, target, left, mid-1)
	} else {
		return search_interval(nums, target, mid+1, right)
	}
}
