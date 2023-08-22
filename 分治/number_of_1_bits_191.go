package divide_conquer

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
func hammingWeight1(num uint32) int {
	var ans int
	for num > 0 {
		ans++
		// n & (n - 1) 的技巧可以消除二进制中最后一个 1
		// 对于任意整数 n，n&(n−1)的结果是将 nnn 的二进制表示的最后一个 1变成 0之后的整数，即 n&(n−1) 的二进制表示中的 1的个数比 n的二进制表示中的 1的个数少 1个
		num &= num - 1
	}
	return ans
}

// 方法二：分治 (左右两部分之和)
// 时间复杂度：O(logN)
// 空间复杂度：O(1)

func hammingWeight(n uint32) int {
	// 每组1个位置，右移1位对齐并删除冗余数据（&01010101）。
	n = (n & 0x55555555) + ((n >> 1) & 0x55555555)
	// 每组2个位置，右移2位对齐并删除冗余数据（&00110011）。
	n = (n & 0x33333333) + ((n >> 2) & 0x33333333)
	// 每组4个位置，右移4位对齐并删除冗余数据（&00001111）。
	n = (n & 0x0f0f0f0f) + ((n >> 4) & 0x0f0f0f0f)
	// 每组8个位置，右移8位对齐并删除冗余数据（&0000000011111111）。
	n = (n & 0x00ff00ff) + ((n >> 8) & 0x00ff00ff)
	// 每组16个位置，右移16位对齐并删除冗余数据（&000000000000000001111111111111111）。
	n = (n & 0x0000ffff) + ((n >> 16) & 0x0000ffff)
	// https://leetcode.cn/problems/number-of-1-bits/solutions/2028759/by-yu-shu-7f-soys/
	return int(n)
}

//颠倒二进制位
//2 的幂
//比特位计数
//二进制手表
//汉明距离
//交替位二进制数
//二进制表示中质数个计算置位
