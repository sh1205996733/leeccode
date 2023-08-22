package binary_search

// https://leetcode.cn/problems/search-a-2d-matrix/
// 方法一：二分查找 矩阵从左到右 从上到下满足递增的性质，所以可以把二维数组看成一个一维递增的数组，然后进行二分查找，只需要将一位坐标转换成二维坐标
// 时间复杂度：O(logNM)
// 空间复杂度：O(1)
func searchMatrix(matrix [][]int, target int) bool {
	m, n := len(matrix), len(matrix[0])
	low, high := 0, m*n-1
	for low <= high {
		mid := (high-low)>>1 + low
		x := matrix[mid/n][mid%n]
		if x < target {
			low = mid + 1
		} else if x > target {
			high = mid - 1
		} else {
			return true
		}
	}
	return false
}
