package backtrack

// 子集
// https://leetcode.cn/problems/subsets/description/
// [1,2,3] --> [1],[2],[3],[2,3],[1,2],[1,3],[1,2,3]
// 方法一：递归
func subsets0(nums []int) [][]int {
	ans := subset(nums, 0)
	return append(ans, []int{})
}

func subset(nums []int, idx int) [][]int {
	if idx == len(nums) {
		return nil
	}
	ans := subset(nums, idx+1)
	for i := range ans {
		ans = append(ans, append([]int{nums[idx]}, ans[i]...))
	}
	ans = append(ans, []int{nums[idx]})
	return ans
}

// 方法二：回溯 [[],[1],[2],[1,2],[3]]
func subsets(nums []int) [][]int {
	var ans = [][]int{{}}
	for _, num := range nums {
		ans = append(ans, []int{num})
		l := len(ans)
		for i := 1; i < l-1; i++ {
			t := append([]int{}, ans[i]...)
			ans = append(ans, append(t, num))
		}
	}
	return ans
}
