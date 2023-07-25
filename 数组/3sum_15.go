package array

import "sort"

// https://leetcode.cn/problems/3sum/
//  三数之和

// 方法一：排序+双指针
// 时间复杂度：O(N)
// 空间复杂度：O(1)
func threeSum(nums []int) [][]int {
	//排序
	sort.Ints(nums)
	var ans [][]int
	for i := 0; i < len(nums); i++ {
		// 和上次遍历一样就直接跳过
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		k := len(nums) - 1
		for j := i + 1; j < len(nums); j++ {
			// 和上次遍历一样就直接跳过
			if j > i+1 && nums[j] == nums[j-1] {
				continue
			}
			for j < k && nums[j]+nums[k] > nums[i]*-1 {
				k--
			}
			if k == j {
				break
			}
			if nums[i]+nums[j]+nums[k] == 0 {
				ans = append(ans, []int{nums[i], nums[j], nums[k]})
			}
		}
	}
	return ans
}
