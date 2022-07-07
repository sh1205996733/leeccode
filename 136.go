/**
 * @Author:shangyc
 * @Date: 2022/5/11 0011
 */

package main

import "fmt"

func main() {
	num := []int{1, 2, 2} //4,1,2,1,2
	fmt.Println(singleNumber(num))
}

//有线性时间复杂度O(n)。 你可以不使用额外空间来实现吗？O(1)
func singleNumber(nums []int) int {
	result := 0
	for i := 0; i < len(nums); i++ {
		result ^= nums[i]
	}
	return result
}
func singleNumber1(nums []int) int {
	m := map[int]uint{}
	for _, num := range nums {
		if _, ok := m[num]; !ok {
			m[num] = 1
		} else {
			m[num]++
		}
	}
	for k, v := range m {
		if v == 1 {
			return k
		}
	}
	return 0
}
