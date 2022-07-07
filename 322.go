package main

import (
	"fmt"
	"math"
)

func main() {
	coins := []int{1, 5, 25, 20}
	fmt.Println(coinChange(coins, 41))
}
func coinChange(coins []int, money int) int {
	if money < 1 || coins == nil || len(coins) == 0 {
		return -1
	}
	dp := make([]int, money+1)
	for i := 1; i <= money; i++ {
		min := math.MaxInt32
		for _, face := range coins {
			if i < face {
				continue
			}
			v := dp[i-face]
			if v < 0 || v >= min {
				continue
			}
			min = v
		}
		if min == math.MaxInt32 {
			dp[i] = -1
		} else {
			dp[i] = min + 1
		}
	}
	fmt.Println(dp)
	return dp[money]
}
func coinChange1(coins []int, money int) int {
	sort(coins)
	idx := len(coins) - 1
	count := 0
	for idx >= 0 {
		for money >= coins[idx] {
			money -= coins[idx]
			fmt.Println(coins[idx])
			count++
		}
		idx--
	}
	if money > 0 {
		return -1
	}
	return count
}
func sort(nums []int) { //冒泡排序
	for end := len(nums) - 1; end > 0; end-- {
		sortedIndex := 1
		for begin := 1; begin <= end; begin++ {
			if nums[begin-1] > nums[begin] {
				nums[begin-1], nums[begin] = nums[begin], nums[begin-1]
				sortedIndex = begin
			}
		}
		end = sortedIndex
	}
}
