package backtrack

// 全排列
// https://leetcode.cn/problems/permutations/

// 方法一：暴力解法
// 时间复杂度：O(N*N!)
// 空间复杂度：O(N!)
func permute0(nums []int) [][]int {
	ret := [][]int{{nums[0]}}
	for i := 1; i < len(nums); i++ {
		ret = permuteN(ret, nums[i])
	}
	return ret
}
func permuteN(nums [][]int, target int) [][]int {
	var ret [][]int
	for i := 0; i < len(nums); i++ {
		t := append([]int{target}, nums[i]...) //A 1 2
		ret = append(ret, append([]int{}, t...))
		for j := 1; j <= len(nums[i]); j++ { //A不断往后移 交换t[A]和前一个t[A-1]的未知
			t[j], t[j-1] = t[j-1], t[j]
			ret = append(ret, append([]int{}, t...))
		}
	}
	return ret
}

// 方法二：回溯思想 + 递归
// 时间复杂度：O(N*N!)
// 空间复杂度：O(N)
func permute(nums []int) [][]int {
	var path []int
	size := 1
	for i := 1; i <= len(nums); i++ {
		size *= i
	}
	ret := make([][]int, 0, size)
	visited := make([]bool, len(nums))
	backTrack(nums, path, visited, 0, &ret)
	return ret
}

func backTrack(nums, path []int, visited []bool, depth int, ret *[][]int) {
	if depth == len(nums) {
		*ret = append(*ret, append([]int{}, path...))
		return
	}
	for i := 0; i < len(nums); i++ { //每次把每个元素都访问一遍
		if visited[i] {
			continue
		}
		// 没有访问过 就加入到path 并将该位置标记为visited,然后进行下一位置递归 最后再将该位置设置false 并移除path
		path = append(path, nums[i])
		visited[i] = true
		backTrack(nums, path, visited, depth+1, ret)
		//回溯 将访问过的元素置为 未访问的
		visited[i] = false
		path = path[:len(path)-1]
	}
}

//40. 组合总和 II
//
//46. 全排列
//
//47. 全排列 II
//
//90. 子集 II
