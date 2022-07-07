package main

import "fmt"

func main() {
	arr := []int{-1, 0, 3, 5, 9, 10, 12}
	fmt.Println(search(arr, 9))
}
func search(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)>>2
		if nums[mid] == target {
			return mid
		} else if nums[mid] > target {
			right = mid - 1
		} else if nums[mid] < target {
			left = mid + 1
		}
	}
	return -1
}
