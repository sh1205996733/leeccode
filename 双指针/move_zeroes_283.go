package pointer

// 移动零
// https://leetcode.cn/problems/move-zeroes/description/

// 双指针
func moveZeroes(nums []int) {
	l := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			if nums[l] == 0 { //减少交换次数
				nums[l], nums[i] = nums[i], nums[l]
			}
			l++
		}
	}
}
