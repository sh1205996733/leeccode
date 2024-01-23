package pointer

import (
	"fmt"
	"testing"
)

func TestMoveZeroes(t *testing.T) {
	nums := []int{0, 1, 0, 3, 12}
	moveZeroes(nums)
	fmt.Println(nums)
}

// 三数之和
func TestThreeSum(t *testing.T) {
	nums := []int{-1, 0, 1, 2, -1, -4}
	ans := threeSum(nums)
	fmt.Println(ans)
}
