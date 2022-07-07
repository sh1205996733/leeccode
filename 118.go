/**
 * @Author:注释中编瞎话
 * @Date: 2022/3/29 0029
 */

package main

import (
	"fmt"
)

func main() {
	d := generate(10)
	fmt.Println(d)
}

//[[1],[1,1],[1,2,1],[1,3,3,1],[1,4,6,4,1]]
//    [1]
//   [1,1]
//  [1,2,1]
// [1,3,3,1]
//[1,4,6,4,1]
//leetcode submit region begin(Prohibit modification and deletion)
func generate(numRows int) [][]int {
	result := make([][]int, numRows)
	for i := 0; i < numRows; i++ {
		result[i] = make([]int, i+1)
		for j := 0; j <= i; j++ {
			if j == 0 || j == i {
				result[i][j] = 1
			} else {
				result[i][j] = result[i-1][j-1] + result[i-1][j]
			}
		}
	}
	return result
}

//leetcode submit region end(Prohibit modification and deletion)
