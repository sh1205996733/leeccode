package pointer

// 乘积小于 K 的子数组
//https://leetcode.cn/problems/subarray-product-less-than-k/description/

// 滑动窗口
// 时间复杂度 O(n)
// 空间复杂度 O(1)
func numSubarrayProductLessThanK(nums []int, k int) int {
	if k < 2 {
		return 0
	}
	var ans int
	left, sum := 0, 1
	for right, num := range nums {
		sum *= num
		for ; sum >= k; left++ {
			sum /= nums[left]
		}
		ans += right - left + 1
	}
	return ans
}
