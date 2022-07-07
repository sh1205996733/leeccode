package main

import "fmt"
import "math"

func main() {
	fmt.Println(mySqrt(65))
}

//数学公式
func mySqrt0(x int) int {
	return int(math.Floor((math.Sqrt(float64(x)))))
}

//二分查找
func mySqrt(x int) int {
	left, right := 0, x
	result := 0
	for left <= right {
		mid := left + (right-left)/2
		if mid*mid <= x {
			left = mid + 1
			result = mid
		} else {
			right = mid - 1
		}
	}
	return result
}
