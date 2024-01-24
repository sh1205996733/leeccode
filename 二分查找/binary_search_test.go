package binary_search

import (
	"fmt"
	"testing"
)

// 二分查找
func TestSearch(t *testing.T) {
	arr := []int{-1, 0, 3, 5, 9, 10, 12}
	fmt.Println(search(arr, 9))
}

// 搜索插入位置
func TestSearchInsert(t *testing.T) {
	fmt.Println(searchInsert([]int{3}, 1))
	fmt.Println(searchInsert([]int{3}, 4))
	fmt.Println(searchInsert([]int{1, 3}, 1))
	fmt.Println(searchInsert([]int{1, 3}, 2))
	fmt.Println(searchInsert([]int{1, 3}, 4))
}

// x 的平方根
func TestMySqrt(t *testing.T) {
	fmt.Println(mySqrt(8))
}

// 最长递增子序列
func TestLengthOfLIS(t *testing.T) {
	nums := []int{10, 9, 2, 5, 1, 7, 101, 18}
	fmt.Println(lengthOfLIS(nums))
}

// 寻找两个正序数组的中位数
func TestFindMedianSortedArrays(t *testing.T) {
	nums := []int{1, 3}
	nums2 := []int{2, 4, 5}
	fmt.Println(findMedianSortedArrays0(nums, nums2))
}

// 寻找峰值
func TestFindPeakElement(t *testing.T) {
	nums := []int{1, 2, 1, 3, 5, 6, 4}
	fmt.Println(findPeakElement(nums))
}

// 搜索二维矩阵
func TestSearchMatrix(t *testing.T) {
	nums := [][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60}, {61, 63, 64, 80}}
	target := 77
	fmt.Println(searchMatrix(nums, target))
}
