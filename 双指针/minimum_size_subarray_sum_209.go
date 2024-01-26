package pointer

// 长度最小的子数组
// https://leetcode.cn/problems/minimum-size-subarray-sum/description/

// 滑动窗口
// 时间复杂度 O(n)
// 空间复杂度 O(1)
func minSubArrayLen(target int, nums []int) int {
	n := len(nums)
	var ans = n + 1
	var left int
	var sum int
	for right, num := range nums {
		sum += num
		//for sum-nums[left] >= target {
		//	sum -= nums[left]
		//	left++
		//}
		//if sum >= target {
		//	ans = min(ans, right-left+1)
		//}
		for sum >= target {
			ans = min(ans, right-left+1)
			sum -= nums[left]
			left++
		}
	}
	if ans <= n {
		return ans
	}
	return 0
}
