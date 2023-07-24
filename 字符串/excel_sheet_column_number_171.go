package string

// Excel 表列序号
// https://leetcode.cn/problems/excel-sheet-column-number/

// 方法一：
// 时间复杂度：O(N)
// 空间复杂度：O(n)

// [A B C D E F G H I J  K  L  M  N  O  P  Q  R  S  T  U  V  W  X  Y  Z]
// [1 2 3 4 5 6 7 8 9 10 11 12 13 14 15 16 17 18 19 20 21 22 23 24 25 26]

// ZY 701 = (701 - Y)/26 =  (26-Z)/26 = 0
// ZZY 18277 = (18277-Y)/26 =  (702 - Z)/26 = (26-Z)/26 = 0
// AAA = 703 = (703 - A)/26 = (27 -A)/26 = (1-A)/26 = 0
func titleToNumber(columnTitle string) int {
	var ans int
	for i := 0; i < len(columnTitle); i++ {
		ans = ans*26 + int(columnTitle[i]-64)
	}
	return ans
}
