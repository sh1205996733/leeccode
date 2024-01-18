package array

// 四数相加 II
// https://leetcode.cn/problems/4sum-ii/description/
func fourSumCount(nums1 []int, nums2 []int, nums3 []int, nums4 []int) int {
	g := make(map[int]int)
	for _, n1 := range nums1 {
		for _, n2 := range nums2 {
			g[n1+n2]++
		}
	}
	var ans int
	for _, n3 := range nums3 {
		for _, n4 := range nums4 {
			ans += g[-n3-n4]
		}
	}
	return ans
}
