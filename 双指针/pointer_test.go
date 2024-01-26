package pointer

import (
	"fmt"
	"testing"
)

func TestMoveZeroes(t *testing.T) {
	nums := []int{0, 1, 0, 3, 12}
	moveZeroes(nums)
	fmt.Println(nums)
}

// 三数之和
func TestThreeSum(t *testing.T) {
	nums := []int{-1, 0, 1, 2, -1, -4}
	ans := threeSum(nums)
	fmt.Println(ans)
}

// 无重复字符的最长子串
func TestLengthOfLongestSubstring(t *testing.T) {
	s := "bb"
	ans := lengthOfLongestSubstring(s)
	fmt.Println(ans)
}

// 长度最小的子数组
func TestMinSubArrayLen(t *testing.T) {
	nums := []int{2, 3, 1, 2, 4, 3}
	target := 70
	ans := minSubArrayLen(target, nums)
	fmt.Println(ans)
}

// 乘积小于 K 的子数组
func TestNumSubarrayProductLessThanK(t *testing.T) {
	nums := []int{10, 5, 2, 6}
	k := 100
	ans := numSubarrayProductLessThanK(nums, k)
	fmt.Println(ans)
}

// 最长重复子数组
func TestFindLength(t *testing.T) {
	nums1 := []int{1, 2, 3, 2, 1}
	nums2 := []int{3, 2, 1, 4, 7}
	ans := findLength(nums1, nums2)
	fmt.Println(ans)
}
