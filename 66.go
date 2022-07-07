package main

import "fmt"

func main() {
	digits := []int{0}
	fmt.Println(plusOne(digits))
}

func plusOne(digits []int) []int {
	index := len(digits) - 1
	for index >= 0 && digits[index] == 9 { //循环
		digits[index] = 0
		index--
	}
	if index == -1 { //999--->1000
		digits = make([]int, len(digits)+1)
		digits[0] = 1
	} else {
		digits[index]++
	}
	return digits
}
