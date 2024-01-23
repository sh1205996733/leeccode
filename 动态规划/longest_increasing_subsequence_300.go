package dp

// 最长递增子序列
// https://leetcode.cn/problems/longest-increasing-subsequence/description/
// 10,9,2,5,3,7,101,18
func lengthOfLIS(nums []int) int {
	if nums == nil || len(nums) == 0 {
		return 0
	}
	size := 0
	heap := make([]int, len(nums))
	for _, num := range nums {
		begin := 0
		end := size
		for begin < end { //遍历所有牌堆的牌顶
			mid := (begin + end) >> 1
			if heap[mid] >= num { //mid左边
				end = mid
			} else {
				begin = mid + 1
			}
		}
		heap[begin] = num
		if begin == size {
			size++
		}

	}
	return size
}

func lengthOfLIS2(nums []int) int {
	if nums == nil || len(nums) == 0 {
		return 0
	}
	dp := make([]int, len(nums))
	dp[0] = 1
	max := 1
	for i := 1; i < len(dp); i++ {
		dp[i] = 1
		for j := 0; j < i; j++ {
			if nums[j] >= nums[i] {
				continue
			}
			if dp[j]+1 > dp[i] {
				dp[i] = dp[j] + 1
			}
		}
		if dp[i] > max {
			max = dp[i]
		}
	}
	return max
}
