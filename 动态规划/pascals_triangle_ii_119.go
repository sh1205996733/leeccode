package dp

// 杨辉三角 II
// https://leetcode.cn/problems/pascals-triangle-ii/description/
// 动态规划
func getRow0(numRows int) []int {
	result := make([]int, numRows+1)
	for i := 0; i <= numRows; i++ {
		temp := append([]int{}, result...)
		for j := 0; j <= i; j++ {
			if j == 0 {
				result[j] = 1
			} else {
				result[j] = temp[j-1] + temp[j]
			}
		}
	}
	return result
}
func getRow1(numRows int) []int {
	result := make([]int, numRows+1)
	for i := 0; i <= numRows; i++ {
		result[0] = 1
		var prev = 1
		for j := 1; j <= i; j++ {
			result[j], prev = prev+result[j], result[j]
		}
	}
	return result
}

func getRow(rowIndex int) []int { //滚动数组
	result := make([]int, rowIndex+1)
	result[0] = 1
	for i := 1; i <= rowIndex; i++ { //第i+1行开始
		for j := i; j > 0; j-- { //第i+1列开始
			result[j] = result[j] + result[j-1]
		}
	}
	return result
}
