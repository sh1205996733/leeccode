package array

// 两个数组的交集 II
// https://leetcode.cn/problems/intersection-of-two-arrays-ii/

// 方法一: 哈希
func intersect(nums1 []int, nums2 []int) []int {
	var ans []int
	if len(nums1) == 0 || len(nums2) == 0 {
		return ans
	}
	if len(nums1) > len(nums2) {
		nums1, nums2 = nums2, nums1
	}
	m := make(map[int]int)
	for _, v := range nums1 {
		m[v]++
	}
	for _, v := range nums2 {
		if m[v] > 0 {
			ans = append(ans, v)
			m[v]--
		}
	}
	return ans
}
