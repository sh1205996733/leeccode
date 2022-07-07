/**
 * @Author:注释中编瞎话
 * @Date: 2022/3/29 0029
 */

package main

import (
	"fmt"
)

func main() {
	d := getRow(6)
	fmt.Println(d)
}

//leetcode submit region begin(Prohibit modification and deletion)
func getRow1(rowIndex int) []int { //动态规划
	result := make([]int, rowIndex+1)
	result[0] = 1
	pre := result[0]
	for i := 0; i <= rowIndex; i++ {
		for j := 1; j <= i; j++ {
			if j == i {
				result[j] = 1
			} else {
				cur := result[j]
				result[j] = pre + cur
				pre = cur
			}
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

//func getRow(rowIndex int) []int {//递归
//	result := make([]int,rowIndex+1)
//	getRowResult(rowIndex,result)
//	return result
//}
//
//func getRowResult(index int, result []int) {
//	if index == 0 {
//		result[index] = 1
//		return
//	}
//	getRowResult(index-1,result)
//	pre := result[0]
//	for j := 1; j <= index; j++ {
//		if j == index {
//			result[j] = 1
//		}else {
//			cur := result[j]
//			result[j] = pre + cur
//			pre = cur
//		}
//	}
//}

//leetcode submit region end(Prohibit modification and deletion)
//1 0 0 0
//1 1 0 0
//1 2 1 0
//1 3 3 1
