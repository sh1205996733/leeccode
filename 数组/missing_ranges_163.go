package array

import "strconv"

// 缺失的区间
// https://leetcode.cn/problems/missing-ranges/description/
// 遍历
func findMissingRanges(nums []int, lower, upper int) []string {
	var ans []string
	var pre = lower - 1
	for i := 0; i < len(nums); i++ {
		if nums[i]-pre == 2 { //说明 pre 和 nums[i] 中间缺失了一个数字，这个缺失的数字就是 pre + 1，也可以用 nums[i] - 1 表示
			ans = append(ans, strconv.Itoa(pre+1))
		} else if nums[i]-pre >= 3 { //说明 pre 和 nums[i] 之间缺失了两个以上的数字，缺失的数字的范围是 pre + 1 到 nums[i] - 1
			ans = append(ans, strconv.Itoa(pre+1)+"->"+strconv.Itoa(nums[i]-1))
		}
		pre = nums[i]
	}
	if upper-pre == 1 {
		ans = append(ans, strconv.Itoa(pre+1))
	} else if upper-pre >= 2 {
		ans = append(ans, strconv.Itoa(pre+1)+"->"+strconv.Itoa(upper))
	}
	return ans
}
