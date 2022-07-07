package main

import "fmt"

func main() {
	nums := []int{1, 3, 5, 6}
	fmt.Println(searchInsert(nums, 0))
}

func searchInsert(nums []int, target int) int {
	begin, end := 0, len(nums)
	for begin < end {
		mid := (begin + end) >> 1
		if nums[mid] > target {
			end = mid
		} else if nums[mid] == target {
			return mid
		} else {
			begin = mid + 1
		}
	}
	return begin
}
