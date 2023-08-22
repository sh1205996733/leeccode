package array

// 合并两个有序数组
// https://leetcode.cn/problems/merge-sorted-array/

// 方法一：遍历
// 时间复杂度：O(n)
// 空间复杂度：O(1)
func merge(nums1 []int, m int, nums2 []int, n int) {
	i := m - 1
	j := n - 1
	cur := len(nums1) - 1
	for j >= 0 {
		if i >= 0 && nums1[i] > nums2[j] {
			nums1[cur] = nums1[i]
			i--
		} else { //nums1[i] <= nums2[j]
			nums1[cur] = nums2[j]
			j--
		}
		cur--
	}
}
