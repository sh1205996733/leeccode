package string

import (
	"strconv"
)

// Fizz Buzz
// https://leetcode.cn/problems/fizz-buzz/description/
func fizzBuzz0(n int) []string {
	var ans []string
	for i := 1; i <= n; i++ {
		if i%3 == 0 && i%5 == 0 {
			ans = append(ans, "FizzBuzz")
		} else if i%3 == 0 {
			ans = append(ans, "Fizz")
		} else if i%5 == 0 {
			ans = append(ans, "Buzz")
		} else {
			ans = append(ans, strconv.Itoa(i))
		}
	}
	return ans
}

// 素数筛选法
func fizzBuzz(n int) []string {
	ans := make([]string, n)
	for i := 0; i < n; i++ {
		ans[i] = strconv.Itoa(i + 1)
	}
	for i := 3; i <= n; i += 3 {
		ans[i-1] = "Fizz"
	}
	for i := 5; i <= n; i += 5 {
		ans[i-1] = "Buzz"
	}
	for i := 15; i <= n; i += 15 {
		ans[i-1] = "FizzBuzz"
	}
	return ans
}
