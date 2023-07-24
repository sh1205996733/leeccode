package bit

import "fmt"

// Excel 表列序号
// https://leetcode.cn/problems/reverse-bits/

// 方法一：循环 num每次向右移动一位 直到num == 0
// 时间复杂度：O(logN)
// 空间复杂度：O(1)

// 00000010100101000001111010011100
// 00111001011110000010100101000000
func reverseBits(num uint32) uint32 {
	var ans uint32
	for i := 0; i < 32 && num > 0; i++ {
		ans |= num & 1 << (31 - i)
		fmt.Println(ans)
		num >>= 1
	}
	return ans
}
