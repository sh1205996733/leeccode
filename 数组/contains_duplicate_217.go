package array

import "sort"

// 存在重复元素
// https://leetcode.cn/problems/contains-duplicate/description/
func containsDuplicate(nums []int) bool {
	sort.Ints(nums)
	for i := 1; i < len(nums); i++ {
		if nums[i] == nums[i-1] {
			return true
		}
	}
	return false
}
func containsDuplicate2(nums []int) bool {
	m := make(map[int]struct{}, 0)
	for i := len(nums) - 1; i >= 0; i-- {
		if _, ok := m[nums[i]]; ok {
			return true
		} else {
			m[nums[i]] = struct{}{}
		}
	}
	return false
}
