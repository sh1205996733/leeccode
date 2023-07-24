package monotonicstack

import (
	"fmt"
	"testing"
)

// 下一个更大元素 I
func TestNextGreaterElement(t *testing.T) {
	num1 := []int{4, 1, 2}
	num2 := []int{1, 3, 4, 2}
	fmt.Println(nextGreaterElement(num1, num2))
}

// 下一个更大元素 II
func TestNextGreaterElements(t *testing.T) {
	nums := []int{1, 1, 1, 1}
	fmt.Println(nextGreaterElements(nums))
}

// 商品折扣后的最终价格
func TestFinalPrices(t *testing.T) {
	nums := []int{1, 1, 1, 1}
	fmt.Println(finalPrices(nums))
}

// 每日温度
func TestDailyTemperatures(t *testing.T) {
	nums := []int{30, 30, 50, 60}
	fmt.Println(dailyTemperatures(nums))
}

// 接雨水
func TestTrap(t *testing.T) {
	nums := []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}
	fmt.Println(trap(nums))
}
