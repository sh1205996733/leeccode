package array

// 丢失的数字
// https://leetcode.cn/problems/missing-number/description/
func missingNumber0(nums []int) int {
	var n = len(nums)
	var ans = (1 + n) * n >> 1
	for i := 0; i < n; i++ {
		ans -= nums[i]
	}
	return ans
}

// 异或
func missingNumber(nums []int) int {
	var xor int
	for i, num := range nums {
		xor ^= i ^ num
	}
	return xor ^ len(nums)
}
