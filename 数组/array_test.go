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
