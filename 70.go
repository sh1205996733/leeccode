package main

import "fmt"

func main() {
	fmt.Println(climbStairs1(45))
	fmt.Println(climbStairs2(45))
}

func climbStairs1(n int) int {
	m := make(map[int]int)
	m[0] = 1
	for i := 1; i <= n; i++ {
		m[i] = m[i-1] + m[i-2]
	}
	return m[n]
}

func climbStairs0(n int) int {
	if n <= 2 {
		return n
	}
	return climbStairs0(n-1) + climbStairs0(n-2)
}

func climbStairs2(n int) int {
	p, q, r := 0, 0, 1
	for i := 1; i <= n; i++ {
		p = q
		q = r
		r = p + q
	}
	return r
}

func climbStairs3(n int) int {
	p, q, r := 0, 0, 1
	for i := 1; i <= n; i++ {
		p = q
		q = r
		r = p + q
	}
	return r
}
