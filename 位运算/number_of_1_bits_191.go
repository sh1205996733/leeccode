package bit

// https://leetcode.cn/problems/number-of-1-bits/
// 位1的个数

// 方法一：位运算
// 时间复杂度：O(N)
// 空间复杂度：O(1)
func hammingWeight0(num uint32) int {
	ans := 0
	for num > 0 {
		ans += int(num & 1)
		num >>= 1
	}
	return ans
}

// 方法一：位运算优化
// 时间复杂度：O(logN)
// 空间复杂度：O(1)
func hammingWeight(num uint32) int {
	var ans int
	for num > 0 {
		ans++
		// n & (n - 1) 的技巧可以消除二进制中最后一个 1
		num &= num - 1
	}
	return ans
}

//颠倒二进制位
//2 的幂
//比特位计数
//二进制手表
//汉明距离
//交替位二进制数
//二进制表示中质数个计算置位
