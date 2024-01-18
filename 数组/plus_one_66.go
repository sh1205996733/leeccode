package array

// 加一
// https://leetcode.cn/problems/plus-one/description/

// [1,2,3]-->[1,2,4] [9,9]-->[1,0,0]
func plusOne(digits []int) []int {
	index := len(digits) - 1
	for index >= 0 && digits[index] == 9 { //循环
		digits[index] = 0
		index--
	}
	if index == -1 { //999--->1000
		digits = make([]int, len(digits)+1)
		digits[0] = 1
	} else {
		digits[index]++
	}
	return digits
}
