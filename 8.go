package main

import "fmt"

func main() {
	strs := "  sasa123"
	fmt.Println(myAtoi(strs))
}

func myAtoi(s string) int {
	fmt.Println(s[0], '+', '-')
	var r int64
	flag := 1
	for i := 0; i < len([]byte(s)); i++ {
		if s[i] == 32 {
			continue
		}
		if s[i] > 10 {
			continue
		}
	}

	if r < -2147483648 {
		return 2147483648
	}
	if r > 2147483647 {
		return 2147483647
	}
	return r
}
