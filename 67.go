/**
 * @Author:shangyc
 * @Date: 2022/5/11 0011
 */

package main

import (
	"fmt"
	"strconv"
)

func main() {
	cases := []struct {
		A, B   string
		Expect string
	}{
		{"1010", "1011", "10101"},
		{"11", "11", "110"},
		{"11", "1", "100"},
		{"0", "1", "1"},
		{"1", "1", "10"},
		{"11", "10111", "11010"},
		{"11", "10011", "10110"},
	}
	for _, c := range cases {
		if !(addBinary(c.A, c.B) == c.Expect) {
			fmt.Println(c.A, c.B, "测试不能通过")
			return
		}
		//fmt.Println(c.A, c.B, addBinary(c.A, c.B) == c.Expect)
	}
	fmt.Println("测试通过")
}

func addBinary(a string, b string) string {
	result := ""
	if len(a) < len(b) {
		a, b = b, a
	}
	start, end := len(a)-1, len(a)-len(b)
	var carry byte //进位
	for i := start; i >= 0; i-- {
		carry = a[i] - '0' + carry
		if i >= end {
			carry += b[i-end] - '0'
		}
		result = strconv.Itoa(int(carry%2)) + result
		carry = carry / 2
	}
	if carry == 1 {
		result = "1" + result
	}
	return result
}

func addBinary1(a string, b string) string {
	result := ""
	if len(a) < len(b) {
		a, b = b, a
	}
	start, end := len(a)-1, len(a)-len(b)
	var carry = "0" //进位
	for i := start; i >= 0; i-- {
		d := a[i] - '0' + carry[0] - '0'
		if i >= end {
			d = d + b[i-end] - '0'
		}
		if d > 1 {
			if d == 2 {
				result = "0" + result
			} else {
				result = "1" + result
			}
			carry = "1"
		} else {
			result = string(d+'0') + result
			carry = "0"
		}
	}
	if carry == "1" {
		result = carry + result
	}
	return result
}
