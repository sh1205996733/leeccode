package math

// 3 的幂
// https://leetcode.cn/problems/power-of-three/description/

// 方法一: 循环
func isPowerOfThree0(n int) bool {
	for n > 0 && n%3 == 0 {
		n /= 3
	}
	return n == 1
}

// 方法二:数学
// 判断是否为最大 333 的幂的约数
// 在题目给定的 32 位有符号整数的范围内，最大的 3 的幂为 3^19=11622614673.我们只需要判断 n 是否是 3^19的约数即可

func isPowerOfThree(n int) bool {
	return n > 0 && 1162261467%n == 0
}
