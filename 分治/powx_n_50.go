package divide_conquer

// Pow(x,n)
// https://leetcode.cn/problems/powx-n/

// 方法一: 分治 (快速幂 + 递归) 从 x 开始，每次直接把上一次的结果进行平方，计算 6 次就可以得到 x^64的值，而不需要对 x 乘 63 次x。
// 时间复杂度：O(logN)
// 空间复杂度：O(logN)
func myPow(x float64, n int) float64 {
	if x == 0 {
		return 0
	}
	if n < 0 {
		return 1 / quickMul(x, -n)
	}
	return quickMul(x, n)
}

func quickMul0(x float64, n int) float64 {
	if n == 0 {
		return 1
	}
	y := quickMul(x, n/2)
	if n&1 == 1 { //奇数
		return y * y * x
	}
	return y * y
}

// 方法二: 二进制 快速幂 + 迭代（去掉递归）
// 时间复杂度：O(logN)
// 空间复杂度：O(1)
func quickMul(x float64, n int) float64 {
	ans := 1.0
	// 贡献的初始值为 x
	base := x
	// 在对 N 进行二进制拆分的同时计算答案
	for n > 0 {
		if n&1 == 1 { //奇数
			ans *= base
		}
		// 将贡献不断地平方
		base *= base
		n >>= 1
	}
	return ans
}
