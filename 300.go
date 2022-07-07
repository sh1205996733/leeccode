package main

import "fmt"

func main() {
	nums := []int{10, 2, 2, 5, 1, 7, 101, 18}
	fmt.Println(lengthOfLIS(nums))
}

func lengthOfLIS(nums []int) int {
	if nums == nil || len(nums) == 0 {
		return 0
	}
	size := 0
	top := make([]int, len(nums))
	for _, num := range nums {
		begin := 0
		end := size
		for begin < end { //遍历所有牌堆的牌顶
			mid := (begin + end) >> 1
			if top[mid] >= num { //mid左边
				end = mid
			} else {
				begin = mid + 1
			}
		}
		top[begin] = num
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
