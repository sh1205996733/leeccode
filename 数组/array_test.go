package array

import (
	"fmt"
	"testing"
)

func TestRotate(t *testing.T) {
	// [[1, 2, 3, 31, 32], [4, 5, 6, 61, 62], [7, 8, 9, 91, 92], [41, 51, 61, 21, 23], [34, 23, 56, 66, 56]]
	matrix := [][]int{{5, 1, 9, 11}, {2, 4, 8, 10}, {13, 3, 6, 7}, {15, 14, 12, 16}}
	//matrix := [][]int{{5, 1}, {2, 4}}
	printRotate(matrix)
	rotate(matrix)
	printRotate(matrix)
}

func printRotate(data [][]int) {
	for _, row := range data {
		for _, col := range row {
			fmt.Printf("%d ", col)
		}
		fmt.Println()
	}
	fmt.Println("---------------")
}

func TestMajorityElement(t *testing.T) {
	nums := []int{1, 1, 2, 2, 3, 2, 2}
	num := majorityElement(nums)
	fmt.Println(num)
}

func TestSortColors(t *testing.T) {
	nums := []int{2, 1, 0}
	sortColors(nums)
	fmt.Println(nums)
}

// 三数之和
func TestThreeSum(t *testing.T) {
	nums := []int{-1, 0, 1, 2, -1, -4}
	ans := threeSum(nums)
	fmt.Println(ans)
}

// 合并两个有序数组
func TestMerge(t *testing.T) {
	nums1 := []int{1, 2, 3, 0, 0, 0}
	m := 3
	nums2 := []int{2, 5, 6}
	n := 3
	merge(nums1, m, nums2, n)
}

func TestFourSum(t *testing.T) {
	nums := []int{1, 0, -1, 0, -2, 2}
	fmt.Println(fourSum(nums, 0))
}

func TestFourSumCount(t *testing.T) {
	nums1, nums2, nums3, nums4 := []int{1, 2}, []int{-2, -1}, []int{-1, 2}, []int{0, 2}
	fmt.Println(fourSumCount(nums1, nums2, nums3, nums4))
}

func TestPlusOne(t *testing.T) {
	digits := []int{0}
	fmt.Println(plusOne(digits))
}

func TestMissingNumber(t *testing.T) {
	digits := []int{9, 6, 4, 2, 3, 5, 7, 0, 1}
	fmt.Println(missingNumber(digits))
}

func TestSingleNumber(t *testing.T) {
	digits := []int{4, 1, 2, 1, 2}
	fmt.Println(singleNumber(digits))
}
