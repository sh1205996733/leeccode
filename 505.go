package main

import "fmt"

func main() {
	fmt.Println(fib5(7))
}

/* 0 1 2 3 4 5
 * 0 1 1 2 3 5 8 13 ....
 */

//递归
func fib1(n int) int {
	if n <= 1 {
		return n
	}
	return fib1(n-1) + fib1(n-2)
}

//循环方式一
func fib2(n int) int {
	if n <= 1 {
		return n
	}
	start, end := 0, 1
	for i := 0; i < n-1; i++ {
		sum := start + end
		start = end
		end = sum
	}
	return end
}

//循环方式二
func fib3(n int) int {
	if n <= 1 {
		return n
	}
	start, end := 0, 1
	for i := 0; i < n-1; i++ {
		end += start
		start = end - start
	}
	return end
}

//循环加数组
func fib4(n int) int {
	if n <= 1 {
		return n
	}
	arr := make([]int, n+1)
	arr[0], arr[1] = 0, 1
	for i := 2; i <= n; i++ {
		arr[i] = arr[i-1] + arr[i-2]
	}
	return arr[n]
}

//循环加map
func fib5(n int) int {
	if n <= 1 {
		return n
	}
	arr := make(map[int]int)
	arr[0], arr[1] = 0, 1
	for i := 2; i <= n; i++ {
		arr[i] = arr[i-1] + arr[i-2]
	}
	return arr[n]
}
