package array

// 旋转图像
// https://leetcode.cn/problems/rotate-image/

// 方法一：暴力法
// 时间复杂度：O(N*M)
// 空间复杂度：O(1)
func rotate(m [][]int) {
	n := len(m)
	for row := 0; row < n/2; row++ {
		for col := 0; col < (n+1)/2; col++ {
			var t = m[row][col]                             // 遍历第一行
			t, m[col][n-1-row] = m[col][n-1-row], t         // 遍历最后一列
			t, m[n-1-row][n-1-col] = m[n-1-row][n-1-col], t // 遍历最后一行
			t, m[n-1-col][row] = m[n-1-col][row], t         // 遍历第一列
			m[row][col] = t
		}
	}
}

/*
	[0,0] [0,1] [0,2]
	[1,0] [1,1] [1,2]
	[2,0] [2,1] [2,2]


[0,0] -> [0,2]、[0,2] -> [2,2]、[2,2] -> [2,0]、[2,0] -> [0,2]
[0,1] -> [1,2]、[1,2] -> [2,1]、[2,1] -> [1,0]、[1,0] -> [0,1]
*/
