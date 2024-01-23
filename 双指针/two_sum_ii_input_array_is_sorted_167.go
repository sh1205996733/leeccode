package pointer

// 两数之和 II - 输入有序数组
// https://leetcode.cn/problems/two-sum-ii-input-array-is-sorted/description/

// 方法一: 相向双指针
func twoSum(numbers []int, target int) []int {
	left, right := 0, len(numbers)-1
	for left < right {
		sum := numbers[left] + numbers[right]
		if sum > target {
			right--
		} else if sum < target {
			left++
		} else {
			return []int{left + 1, right + 1}
		}
	}
	return []int{}
}
